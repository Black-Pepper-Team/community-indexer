package core

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"gitlab.com/distributed_lab/logan/v3"

	"github.com/black-pepper-team/community-indexer/internal/config"
	"github.com/black-pepper-team/community-indexer/internal/data"
	"github.com/black-pepper-team/community-indexer/internal/data/postgres"
)

type Core struct {
	log *logan.Entry
	db  data.MasterQ

	ethClient  *ethclient.Client
	privateKey *ecdsa.PrivateKey
	chainID    *big.Int
	address    common.Address
}

func New(ctx context.Context, cfg config.Config) (*Core, error) {
	chainID, err := cfg.EthClient().EthClient.ChainID(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get chainID: %w", err)
	}

	return &Core{
		log:        cfg.Log(),
		db:         postgres.NewMasterQ(cfg.DB()),
		ethClient:  cfg.EthClient().EthClient,
		privateKey: cfg.EthClient().PrivateKey,
		address:    crypto.PubkeyToAddress(cfg.EthClient().PrivateKey.PublicKey),
		chainID:    chainID,
	}, nil
}

func (c *Core) GetCommunitiesList() ([]data.Community, error) {
	communities, err := c.db.New().CommunitiesQ().Select()
	if err != nil {
		c.log.WithError(err).Error("Failed to get communities list")
		return nil, err
	}

	return communities, nil
}

func (c *Core) CreateCommunity(ownerAddress, collectionName, collectionSymbol string) (int64, error) {

	return communityId, nil
}
