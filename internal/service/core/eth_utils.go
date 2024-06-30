package core

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	// invalidSignatureError Keccak256("Error(string)")[:4]
	invalidSignatureError = []byte{0x08, 0xc3, 0x79, 0xa0}

	abiString, _ = abi.NewType("string", "", nil)
)

var ErrTransactionFailed = errors.New("transaction failed")

func (c *Core) waitMined(ctx context.Context, tx *types.Transaction) (*types.Receipt, error) {
	receipt, err := bind.WaitMined(ctx, c.ethClient, tx)
	if err != nil {
		return nil, fmt.Errorf("failed to get mined tx: %w", err)
	}

	if receipt.Status == types.ReceiptStatusFailed {
		reason, err := checkTxErrorReason(ctx, tx, c.ethClient, c.address)
		if err != nil {
			return nil, fmt.Errorf("failed to get tx error reason: %w", err)
		}
		return nil, fmt.Errorf("%w: %s", ErrTransactionFailed, reason)
	}

	return receipt, nil
}

func (c *Core) retryChainCall(
	ctx context.Context,
	privKey *ecdsa.PrivateKey,
	contractCall func(signer *bind.TransactOpts) error,
) error {
	signer, err := c.newSigner(ctx, privKey)
	if err != nil {
		return fmt.Errorf("failed to get new signer: %w", err)
	}

	for {
		err := contractCall(signer)
		if err == nil {
			break
		}
		if !errors.Is(err, core.ErrNonceTooLow) {
			return fmt.Errorf("failed to call contract: %w", err)
		}

		signer.Nonce.Add(signer.Nonce, big.NewInt(1))
	}
	return nil
}

func (c *Core) newSigner(ctx context.Context, privKey *ecdsa.PrivateKey) (*bind.TransactOpts, error) {
	nonce, err := c.ethClient.PendingNonceAt(ctx, crypto.PubkeyToAddress(privKey.PublicKey))
	if err != nil {
		return nil, fmt.Errorf("failed to get nonce: %w", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privKey, c.chainID)
	if err != nil {
		return nil, fmt.Errorf("failed to create transaction signer: %w", err)
	}
	auth.Nonce = big.NewInt(int64(nonce))

	return auth, nil
}

func checkTxErrorReason(
	ctx context.Context,
	tx *types.Transaction,
	ethClient *ethclient.Client,
	addr common.Address,
) (string, error) {
	msg := ethereum.CallMsg{
		From:     addr,
		To:       tx.To(),
		Gas:      tx.Gas(),
		GasPrice: tx.GasPrice(),
		Value:    tx.Value(),
		Data:     tx.Data(),
	}

	res, err := ethClient.CallContract(ctx, msg, nil)
	if err != nil {
		return "", fmt.Errorf("failed to CallContract: %w", err)
	}

	return unpackError(res)
}

func unpackError(result []byte) (string, error) {
	if len(result) < 4 || !bytes.Equal(result[:4], invalidSignatureError) {
		return "tx result is not Error(string)", fmt.Errorf("TX result not of type Error(string)")
	}

	values, err := abi.Arguments{{Type: abiString}}.UnpackValues(result[4:])
	if err != nil {
		return "invalid tx result", fmt.Errorf("unpacking revert reason: %w", err)
	}

	return values[0].(string), nil
}
