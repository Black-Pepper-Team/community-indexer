package requests

import (
	"crypto/ecdsa"
	"encoding/json"
	"net/http"

	"github.com/ethereum/go-ethereum/crypto"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type createCommunityRequest struct {
	CollectionName   string `json:"collection_name"`
	CollectionSymbol string `json:"collection_symbol"`
	PrivateKey       string `json:"private_key"`
}

type CreateCommunityRequest struct {
	CollectionName   string
	CollectionSymbol string
	PrivateKey       *ecdsa.PrivateKey
}

func NewCreateCommunity(r *http.Request) (*CreateCommunityRequest, error) {
	var requestBody createCommunityRequest

	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		return nil, errors.Wrap(err, "failed to decode json request body")
	}

	if err := requestBody.validate(); err != nil {
		return nil, err
	}

	return requestBody.parse(), nil
}

// nolint
func (r *createCommunityRequest) validate() error {
	return validation.Errors{
		"body/private_key": validation.Validate(
			r.PrivateKey, validation.Required, validation.Length(64, 64),
		),
		"body/collection_name": validation.Validate(
			r.CollectionName, validation.Required,
		),
		"body/collection_symbol": validation.Validate(
			r.CollectionSymbol, validation.Required,
		),
	}.Filter()
}

func (r *createCommunityRequest) parse() *CreateCommunityRequest {
	sk, _ := crypto.HexToECDSA(r.PrivateKey)
	return &CreateCommunityRequest{
		PrivateKey:       sk,
		CollectionName:   r.CollectionName,
		CollectionSymbol: r.CollectionSymbol,
	}
}
