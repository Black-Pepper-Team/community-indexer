package requests

import (
	"encoding/hex"
	"encoding/json"
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/iden3/go-iden3-crypto/babyjub"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type sendMessageRequest struct {
	NFTID         string `json:"nft_id"`
	NFTOwner      string `json:"nft_owner"`
	ContractId    string `json:"contract_id"`
	BJJPrivateKey string `json:"bjj_private_key"`
	Message       string `json:"message"`
}

type SendMessageRequest struct {
	NFTID         *big.Int           `json:"nft_id"`
	NFTOwner      common.Address     `json:"nft_owner"`
	ContractId    common.Address     `json:"contract_id"`
	BJJPrivateKey babyjub.PrivateKey `json:"bjj_private_key"`
	Message       string             `json:"message"`
}

func NewSendMessage(r *http.Request) (*SendMessageRequest, error) {
	var requestBody sendMessageRequest

	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		return nil, errors.Wrap(err, "failed to decode json request body")
	}

	if err := requestBody.validate(); err != nil {
		return nil, err
	}

	return requestBody.parse(), nil
}

// nolint
func (r *sendMessageRequest) validate() error {
	return validation.Errors{
		"body/nft_owner": validation.Validate(
			r.NFTOwner, validation.Required, validation.By(MustBeValidAddress),
		),
		"body/contract_id": validation.Validate(
			r.ContractId, validation.Required, validation.By(MustBeValidAddress),
		),
	}.Filter()
}

func (r *sendMessageRequest) parse() *SendMessageRequest {
	bjjSkBytes, _ := hex.DecodeString(r.BJJPrivateKey)
	bjjPrivateKey := babyjub.PrivateKey(bjjSkBytes)
	nftId, _ := new(big.Int).SetString(r.NFTID, 10)

	return &SendMessageRequest{
		NFTID:         nftId,
		NFTOwner:      common.HexToAddress(r.NFTOwner),
		ContractId:    common.HexToAddress(r.ContractId),
		BJJPrivateKey: bjjPrivateKey,
		Message:       r.Message,
	}
}
