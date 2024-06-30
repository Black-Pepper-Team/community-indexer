package core

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/google/uuid"
	"gitlab.com/distributed_lab/logan/v3"

	"github.com/black-pepper-team/community-indexer/contracts"
	"github.com/black-pepper-team/community-indexer/internal/config"
	"github.com/black-pepper-team/community-indexer/internal/data"
	"github.com/black-pepper-team/community-indexer/internal/data/postgres"
)

const (
	VerifiableCommitmentCircuitName = "VerifiableCommitment"
	PostMessageCircuitName          = "PostMessage"
)

type Core struct {
	log *logan.Entry
	db  data.MasterQ

	ctx context.Context

	ethClient  *ethclient.Client
	privateKey *ecdsa.PrivateKey
	chainID    *big.Int
	address    common.Address

	authStorageContract *contracts.AuthenticationStorage
	chatContract        *contracts.Chat
	registerRequests    map[uuid.UUID]RegisterRequest
	registeredUsers     RegisterStorage

	circuits Circuits
}

func New(ctx context.Context, cfg config.Config) (*Core, error) {
	chainID, err := cfg.EthClient().EthClient.ChainID(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get chainID: %w", err)
	}

	circuits, err := readCircuits(cfg.Circom().PathToCircuits)
	if err != nil {
		return nil, fmt.Errorf("failed to read state transition circuits: %w", err)
	}

	authStorageContract, err := contracts.NewAuthenticationStorage(
		cfg.EthClient().AuthStorageContract,
		cfg.EthClient().EthClient,
	)

	chatContract, err := contracts.NewChat(
		cfg.EthClient().ChatContract,
		cfg.EthClient().EthClient,
	)

	return &Core{
		log:                 cfg.Log(),
		ctx:                 ctx,
		db:                  postgres.NewMasterQ(cfg.DB()),
		ethClient:           cfg.EthClient().EthClient,
		privateKey:          cfg.EthClient().PrivateKey,
		address:             crypto.PubkeyToAddress(cfg.EthClient().PrivateKey.PublicKey),
		chainID:             chainID,
		circuits:            circuits,
		chatContract:        chatContract,
		authStorageContract: authStorageContract,
		registeredUsers:     make(RegisterStorage),
	}, nil
}

func readCircuits(pathToCircuits string) (Circuits, error) {
	var err error
	circuits := make(Circuits)

	circuits[VerifiableCommitmentCircuitName] = make(map[circuitKey][]byte)
	circuits[VerifiableCommitmentCircuitName][ZKEY], err = ReadFileByPath(pathToCircuits, fmt.Sprintf("%s.zkey", VerifiableCommitmentCircuitName))
	if err != nil {
		return nil, fmt.Errorf("failed to read state transition circuit final: %w", err)
	}

	circuits[PostMessageCircuitName] = make(map[circuitKey][]byte)
	circuits[PostMessageCircuitName][ZKEY], err = ReadFileByPath(pathToCircuits, fmt.Sprintf("%s.zkey", PostMessageCircuitName))
	if err != nil {
		return nil, fmt.Errorf("failed to read state transition circuit wasm: %w", err)
	}

	circuits[VerifiableCommitmentCircuitName][WASM], err = ReadFileByPath(pathToCircuits, fmt.Sprintf("%s.wasm", VerifiableCommitmentCircuitName))
	if err != nil {
		return nil, fmt.Errorf("failed to read state transition circuit final: %w", err)
	}

	circuits[PostMessageCircuitName][WASM], err = ReadFileByPath(pathToCircuits, fmt.Sprintf("%s.wasm", PostMessageCircuitName))
	if err != nil {
		return nil, fmt.Errorf("failed to read state transition circuit wasm: %w", err)
	}

	return circuits, nil
}
