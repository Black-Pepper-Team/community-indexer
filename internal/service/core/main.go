package core

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/google/uuid"
	"gitlab.com/distributed_lab/logan/v3"

	"github.com/black-pepper-team/community-indexer/contracts"
	"github.com/black-pepper-team/community-indexer/internal/config"
	"github.com/black-pepper-team/community-indexer/internal/data"
	"github.com/black-pepper-team/community-indexer/internal/data/postgres"
)

type Core struct {
	log *logan.Entry
	db  data.MasterQ

	ctx context.Context

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
		ctx:        ctx,
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

func (c *Core) CreateCommunity(
	collectionName string,
	collectionSymbol string,
) (*data.Community, error) {
	var err error
	var tx *types.Transaction
	var address common.Address
	err = c.retryChainCall(c.ctx, func(signer *bind.TransactOpts) error {
		signer, err := c.newSigner(c.ctx)
		if err != nil {
			return fmt.Errorf("failed to get new signer: %w", err)
		}

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
		ContractAddress: address,
		OwnerAddress:    c.address,
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
