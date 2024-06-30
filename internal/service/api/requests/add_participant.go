package requests

import (
	"crypto/ecdsa"
	"encoding/json"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type addCommunityParticipantRequest struct {
	PrivateKey         string `json:"private_key"`
	ParticipantAddress string `json:"participant_address"`
	ContractAddress    string `json:"contract_address"`
}

type AddCommunityParticipantRequest struct {
	PrivateKey         *ecdsa.PrivateKey
	ParticipantAddress common.Address
	ContractAddress    common.Address
}

func NewAddCommunityParticipant(r *http.Request) (*AddCommunityParticipantRequest, error) {
	var requestBody addCommunityParticipantRequest

	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		return nil, errors.Wrap(err, "failed to decode json request body")
	}

	if err := requestBody.validate(); err != nil {
		return nil, err
	}

	return requestBody.parse(), nil
}

// nolint
func (r *addCommunityParticipantRequest) validate() error {
	return validation.Errors{
		"body/private_key": validation.Validate(
			r.PrivateKey, validation.Required, validation.Length(64, 64),
		),
		"body/participant_address": validation.Validate(
			r.ParticipantAddress, validation.Required, validation.By(MustBeValidAddress),
		),
		"body/contract_address": validation.Validate(
			r.ContractAddress, validation.Required, validation.By(MustBeValidAddress),
		),
	}.Filter()
}

func (r *addCommunityParticipantRequest) parse() *AddCommunityParticipantRequest {
	sk, _ := crypto.HexToECDSA(r.PrivateKey)
	return &AddCommunityParticipantRequest{
		PrivateKey:         sk,
		ParticipantAddress: common.HexToAddress(r.ParticipantAddress),
		ContractAddress:    common.HexToAddress(r.ContractAddress),
	}
}
