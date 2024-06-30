package core

import (
	"crypto/ecdsa"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/google/uuid"
	"github.com/iden3/go-iden3-crypto/babyjub"
	"github.com/iden3/go-rapidsnark/prover"
	"github.com/iden3/go-rapidsnark/witness/v2"
	"github.com/iden3/go-rapidsnark/witness/wazero"

	"github.com/black-pepper-team/community-indexer/contracts"
	"github.com/black-pepper-team/community-indexer/internal/data"
)

var (
	ErrContractNotFound        = fmt.Errorf("contract not found")
	ErrCreateWitnessCalculator = errors.New("failed to create witness calculator")
)

func (c *Core) GetCommunitiesList() ([]data.Community, error) {
	communities, err := c.db.New().CommunitiesQ().Select()
	if err != nil {
		c.log.WithError(err).Error("Failed to get communities list")
		return nil, err
	}

	return communities, nil
}

func (c *Core) GetCommunityById(communityId uuid.UUID) (*data.Community, error) {
	community, err := c.db.New().CommunitiesQ().WhereID(communityId).Get()
	if err != nil {
		c.log.WithError(err).Error("Failed to get community by ID")
		return nil, err
	}

	if community == nil {
		return nil, nil
	}

	return community, nil
}

func (c *Core) CreateCommunity(
	collectionName string,
	collectionSymbol string,
	privKey *ecdsa.PrivateKey,
) (*data.Community, error) {
	var err error
	var tx *types.Transaction
	var address common.Address
	err = c.retryChainCall(c.ctx, privKey, func(signer *bind.TransactOpts) error {
		address, tx, _, err = contracts.DeployERC721Mock(signer, c.ethClient, collectionName, collectionSymbol)
		if err != nil {
			return fmt.Errorf("failed to deploy new community: %w", err)
		}

		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to call retryChainCall: %w", err)
	}

	communityId, err := uuid.NewRandom()
	if err != nil {
		return nil, fmt.Errorf("failed to generate new UUID: %w", err)
	}

	newCommunity := data.Community{
		ID:              communityId,
		Status:          data.Deploying,
		Name:            collectionName,
		Symbol:          collectionSymbol,
		ContractAddress: address,
		OwnerAddress:    crypto.PubkeyToAddress(privKey.PublicKey),
	}

	if err = c.db.New().CommunitiesQ().Insert(&newCommunity); err != nil {
		return nil, fmt.Errorf("failed to insert new community: %w", err)
	}

	go func() {
		nextStatus := data.Ready

		if _, err := c.waitMined(c.ctx, tx); err != nil {
			c.log.
				WithField("tx-hash", tx.Hash().Hex()).
				WithField("reason", err).
				Errorf("failed to wait tx mined")

			nextStatus = data.FailedToDeploy
		}

		newCommunity.Status = nextStatus
		if newErr := c.db.New().CommunitiesQ().Update(&newCommunity); newErr != nil {
			c.log.
				WithField("reason", newErr.Error()).
				Errorf("failed to update the community status to failed community")
		}

		return
	}()

	return &newCommunity, nil
}

func (c *Core) ImportCommunity(contractAddress common.Address) (*data.Community, error) {
	existedCommunity, err := c.db.CommunitiesQ().WhereContractAddress(contractAddress).Get()
	if err != nil {
		return nil, fmt.Errorf("failed to get community by contract address: %w", err)
	}

	if existedCommunity != nil {
		return existedCommunity, nil
	}

	collectionContract, err := contracts.NewERC721Mock(contractAddress, c.ethClient)
	if err != nil {
		return nil, fmt.Errorf("failed to get ERC721 contract: %w", err)
	}

	collectionName, err := collectionContract.Name(nil)
	if err != nil {
		// Assume that the contract doesn't exists
		return nil, ErrContractNotFound
	}

	collectionSymbol, err := collectionContract.Symbol(nil)
	if err != nil {
		// Assume that the contract doesn't exists
		return nil, ErrContractNotFound
	}

	contractOwner, err := collectionContract.Owner(nil)
	if err != nil {
		// Assume that the contract doesn't exists
		return nil, ErrContractNotFound
	}

	communityId, err := uuid.NewRandom()
	if err != nil {
		return nil, fmt.Errorf("failed to generate new UUID: %w", err)
	}

	newCommunity := &data.Community{
		ID:              communityId,
		Status:          data.Ready,
		Name:            collectionName,
		Symbol:          collectionSymbol,
		ContractAddress: contractAddress,
		OwnerAddress:    contractOwner,
	}

	if err = c.db.New().CommunitiesQ().Insert(newCommunity); err != nil {
		return nil, fmt.Errorf("failed to insert new community: %w", err)
	}

	return newCommunity, nil
}

func (c *Core) RegisterInCommunity(
	bjjPublicKey babyjub.PublicKey,
	nftOwner common.Address,
	contractId common.Address,
	nftId big.Int,
) (*RegisterRequest, error) {
	if c.registeredUsers[nftOwner] != nil {
		registerId := c.registeredUsers[nftOwner][contractId]
		if registerId != nil {
			registryRequest := c.registerRequests[*registerId]

			if registryRequest.Status == Processing || registryRequest.Status == Registered {
				return registryRequest, nil
			}
		}
	}

	wtnsCalculator, err := witness.NewCalculator(
		c.circuits[VerifiableCommitmentCircuitName][WASM],
		witness.WithWasmEngine(wazero.NewCircom2WZWitnessCalculator),
	)
	if err != nil {
		return nil, ErrCreateWitnessCalculator
	}

	timestamp := time.Now().Add(96 * time.Hour)
	deadline := timestamp.Add(300 * time.Second).Unix()
	transitionInputs := VarifiableCommitmentInputs{
		ContractId:     contractId.String(),
		NftId:          nftId.String(),
		NftOwner:       nftOwner.String(),
		Deadline:       strconv.FormatInt(deadline, 10),
		Timestamp:      strconv.FormatInt(timestamp.Unix(), 10),
		BabyJubJubPKAx: bjjPublicKey.X.String(),
		BabyJubJubPKAy: bjjPublicKey.Y.String(),
	}

	transitionInputsRaw, err := json.Marshal(transitionInputs)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal transition inputs: %w", err)
	}

	parsedInputs, err := witness.ParseInputs(transitionInputsRaw)
	if err != nil {
		return nil, fmt.Errorf("failed to parse witness inputs: %w", err)
	}

	wtnsBytes, err := wtnsCalculator.CalculateWTNSBin(parsedInputs, true)
	if err != nil {
		return nil, fmt.Errorf("failed to calculate witnesses: %w", err)
	}

	rapidProof, err := prover.Groth16Prover(c.circuits[VerifiableCommitmentCircuitName][ZKEY], wtnsBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to generate prov with groth16: %w", err)
	}

	contractZKP, err := parseZKPArgs(&ZKProof{
		A:        rapidProof.Proof.A,
		B:        rapidProof.Proof.B,
		C:        rapidProof.Proof.C,
		Protocol: rapidProof.Proof.Protocol,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to parse ZKProof to contract readable: %w", err)
	}

	credentialIdBigInt, ok := new(big.Int).SetString(rapidProof.PubSignals[0], 10)
	if !ok {
		return nil, errors.New("failed to parse credential ID")
	}

	var tx *types.Transaction
	err = c.retryChainCall(c.ctx, c.privateKey, func(signer *bind.TransactOpts) error {
		tx, err = c.authStorageContract.Register(
			signer,
			contractId,
			&nftId,
			nftOwner,
			[32]byte(credentialIdBigInt.Bytes()),
			big.NewInt(deadline),
			*contractZKP,
		)
		if err != nil {
			return fmt.Errorf("failed to call Register: %w", err)
		}

		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to call retryChainCall: %w", err)
	}

	requestId := uuid.New()
	registerRequest := RegisterRequest{
		Id:     requestId,
		Status: Processing,
	}

	c.registerRequests[requestId] = &registerRequest

	if c.registeredUsers[nftOwner] == nil {
		c.registeredUsers[nftOwner] = make(map[common.Address]*uuid.UUID)
	}
	c.registeredUsers[nftOwner][contractId] = &requestId

	go func() {
		nextStatus := Registered

		if _, err := c.waitMined(c.ctx, tx); err != nil {
			c.log.
				WithField("tx-hash", tx.Hash().Hex()).
				WithField("reason", err).
				Errorf("failed to wait tx mined")

			nextStatus = FailedRegister
		}

		registerRequest.Status = nextStatus
		c.registerRequests[requestId] = &registerRequest

		return
	}()

	return &registerRequest, nil
}

type VarifiableCommitmentInputs struct {
	ContractId     string `json:"contractId"`
	NftId          string `json:"nftId"`
	NftOwner       string `json:"nftOwner"`
	Deadline       string `json:"deadline"`
	BabyJubJubPKAx string `json:"babyJubJubPK_Ax"`
	BabyJubJubPKAy string `json:"babyJubJubPK_Ay"`
	Timestamp      string `json:"timestamp"`
}

func (c *Core) GetRegister(registerRequestId uuid.UUID) (*RegisterRequest, error) {
	registerRequest := c.registerRequests[registerRequestId]

	if registerRequest == nil {
		return nil, nil
	}

	return registerRequest, nil
}

func (c *Core) AddCommunityParticipant(
	privKey *ecdsa.PrivateKey,
	participantAddress common.Address,
	contractAddress common.Address,
) (*RegisterRequest, error) {
	community, err := c.db.CommunitiesQ().WhereContractAddress(contractAddress).Get()
	if err != nil {
		return nil, fmt.Errorf("failed to get community by contract address: %w", err)
	}

	if community == nil {
		return nil, ErrContractNotFound
	}

	collectionContract, err := contracts.NewERC721Mock(contractAddress, c.ethClient)
	if err != nil {
		return nil, fmt.Errorf("failed to get ERC721 contract: %w", err)
	}

	var tx *types.Transaction
	err = c.retryChainCall(c.ctx, privKey, func(signer *bind.TransactOpts) error {
		newRandId := babyjub.NewRandPrivKey()
		tx, err = collectionContract.Mint(signer, participantAddress, (newRandId).Scalar().BigInt())
		if err != nil {
			return fmt.Errorf("failed to call Register: %w", err)
		}

		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to call retryChainCall: %w", err)
	}

	requestId := uuid.New()
	mintRequest := RegisterRequest{
		Id:     requestId,
		Status: Processing,
	}

	c.registerRequests[requestId] = &mintRequest

	go func() {
		nextStatus := Registered

		if _, err := c.waitMined(c.ctx, tx); err != nil {
			c.log.
				WithField("tx-hash", tx.Hash().Hex()).
				WithField("reason", err).
				Errorf("failed to wait tx mined")

			nextStatus = FailedRegister
		}

		mintRequest.Status = nextStatus
		c.registerRequests[requestId] = &mintRequest

		return
	}()

	return &mintRequest, nil
}
