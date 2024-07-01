package requests

import (
	"crypto/ecdsa"
	"encoding/json"
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/iden3/go-iden3-crypto/babyjub"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type registerInCommunityRequest struct {
	NFTID        string `json:"nft_id"`
	NFTOwner     string `json:"nft_owner"`
	ContractId   string `json:"contract_id"`
	BJJPublicKey string `json:"bjj_public_key"`
	PrivateKey   string `json:"private_key"`
}

type RegisterInCommunityRequest struct {
	NFTID        big.Int
	NFTOwner     common.Address
	ContractId   common.Address
	BJJPublicKey babyjub.PublicKey
	PrivateKey   *ecdsa.PrivateKey
}

func NewRegisterInCommunity(r *http.Request) (*RegisterInCommunityRequest, error) {
	var requestBody registerInCommunityRequest

	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		return nil, errors.Wrap(err, "failed to decode json request body")
	}

	if err := requestBody.validate(); err != nil {
		return nil, err
	}

	return requestBody.parse(), nil
}

// nolint
func (r *registerInCommunityRequest) validate() error {
	return validation.Errors{
		"body/nft_owner": validation.Validate(
			r.NFTOwner, validation.Required, validation.By(MustBeValidAddress),
		),
		"body/contract_id": validation.Validate(
			r.ContractId, validation.Required, validation.By(MustBeValidAddress),
		),
		"body/bjj_public_key": validation.Validate(
			r.BJJPublicKey, validation.Required, validation.By(MustBeValidBJJPublicKey),
		),
	}.Filter()
}

func (r *registerInCommunityRequest) parse() *RegisterInCommunityRequest {
	var bjjPublicKey babyjub.PublicKey
	(&bjjPublicKey).UnmarshalText([]byte(r.BJJPublicKey))

	nftId, _ := new(big.Int).SetString(r.NFTID, 10)
	sk, _ := crypto.HexToECDSA(r.PrivateKey)

	return &RegisterInCommunityRequest{
		NFTID:        *nftId,
		NFTOwner:     common.HexToAddress(r.NFTOwner),
		ContractId:   common.HexToAddress(r.ContractId),
		BJJPublicKey: bjjPublicKey,
		PrivateKey:   sk,
	}
}

func MustBeValidBJJPublicKey(value interface{}) error {
	valueStr, ok := value.(string)
	if !ok {
		return errors.New("invalid BJJ public key")
	}

	if err := new(babyjub.PublicKey).UnmarshalText([]byte(valueStr)); err != nil {
		return errors.Wrap(err, "invalid BJJ public key")
	}

	return nil
}
