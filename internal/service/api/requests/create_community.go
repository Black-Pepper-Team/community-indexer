package requests

import (
	"encoding/json"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type createCommunityRequest struct {
	CollectionName   string `json:"collection_name"`
	CollectionSymbol string `json:"collection_symbol"`
}

type CreateCommunityRequest struct {
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
		CollectionName:   r.CollectionName,
		CollectionSymbol: r.CollectionSymbol,
	}
}
