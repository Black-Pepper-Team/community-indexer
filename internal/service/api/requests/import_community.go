package requests

import (
	"encoding/json"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type importCommunityRequest struct {
	ContractAddress string `json:"contract_address"`
}

type ImportCommunityRequest struct {
	ContractAddress common.Address `json:"contract_address"`
}

func NewImportCommunity(r *http.Request) (*ImportCommunityRequest, error) {
	var requestBody importCommunityRequest

	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		return nil, errors.Wrap(err, "failed to decode json request body")
	}

	if err := requestBody.validate(); err != nil {
		return nil, err
	}

	return requestBody.parse(), nil
}

// nolint
func (r *importCommunityRequest) validate() error {
	return validation.Errors{
		"body/contract_address": validation.Validate(
			r.ContractAddress, validation.Required, validation.By(MustBeValidAddress),
		),
	}.Filter()
}

func (r *importCommunityRequest) parse() *ImportCommunityRequest {
	return &ImportCommunityRequest{
		ContractAddress: common.HexToAddress(r.ContractAddress),
	}
}

func MustBeValidAddress(src interface{}) error {
	addressRaw, ok := src.(string)
	if !ok {
		return errors.New("it is not a string")
	}

	if !common.IsHexAddress(addressRaw) {
		return errors.New("it is not a valid address")
	}

	return nil
}
