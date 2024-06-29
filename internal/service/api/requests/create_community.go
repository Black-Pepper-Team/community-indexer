package requests

import (
	"encoding/json"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type createCommunityRequest struct {
	OwnerAddress     string `json:"owner_address"`
	CollectionName   string `json:"collection_name"`
	CollectionSymbol string `json:"collection_symbol"`
}

type CreateCommunityRequest struct {
	OwnerAddress     common.Address
	CollectionName   string
	CollectionSymbol string
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
		"body/owner_address": validation.Validate(
			r.OwnerAddress, validation.Required, validation.By(MustBeValidEthAddress),
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
	return &CreateCommunityRequest{
		OwnerAddress:     common.HexToAddress(r.OwnerAddress),
		CollectionName:   r.CollectionName,
		CollectionSymbol: r.CollectionSymbol,
	}
}

func MustBeValidEthAddress(src interface{}) error {
	addressRaw, ok := src.(string)
	if !ok {
		return errors.New("it is not a hash")
	}

	if !common.IsHexAddress(addressRaw) {
		return errors.New("it is not valid ethereum address")
	}

	return nil
}
