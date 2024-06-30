package config

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type EthClientConfig struct {
	EthClient           *ethclient.Client
	PrivateKey          *ecdsa.PrivateKey
	AuthStorageContract common.Address
	ChatContract        common.Address
}

type ethClientConfigRaw struct {
	RpcURL              string `fig:"rpc_url,required"` //nolint
	PrivateKey          string `fig:"private_key,required"`
	AuthStorageContract string `fig:"auth_storage_contract,required"`
	ChatContract        string `fig:"chat_contract,required"`
}

func (c *config) EthClient() *EthClientConfig {
	return c.ethClient.Do(func() interface{} {
		configRaw := ethClientConfigRaw{}
		err := figure.
			Out(&configRaw).
			From(kv.MustGetStringMap(c.getter, "ethereum")).
			Please()
		if err != nil {
			panic(errors.Wrap(err, "failed to figure out ethereum config"))
		}

		ethClient, err := ethclient.Dial(configRaw.RpcURL)
		if err != nil {
			panic(errors.Wrap(err, "failed to create an ethereum client"))
		}

		privateKey, err := crypto.HexToECDSA(configRaw.PrivateKey)
		if err != nil {
			panic(errors.Wrap(err, "failed to parse private key"))
		}

		return &EthClientConfig{
			EthClient:           ethClient,
			PrivateKey:          privateKey,
			AuthStorageContract: common.HexToAddress(configRaw.AuthStorageContract),
			ChatContract:        common.HexToAddress(configRaw.ChatContract),
		}
	}).(*EthClientConfig)
}
