package requests

import (
	"net/http"

	"github.com/go-chi/chi"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/uuid"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

const (
	communityIdPathParam = "community-id"
)

type getCommunityRequest struct {
	CommunityID string
}

type GetCommunityRequest struct {
	CommunityID uuid.UUID
}

func NewGetCommunity(r *http.Request) (*GetCommunityRequest, error) {
	requestRaw := getCommunityRequest{
		CommunityID: chi.URLParam(r, communityIdPathParam),
	}

	if err := requestRaw.validate(); err != nil {
		return nil, err
	}

	return requestRaw.parse(), nil
}

// nolint
func (r *getCommunityRequest) validate() error {
	return validation.Errors{
		"path/{revocation-id}": validation.Validate(
			r.CommunityID, validation.Required, validation.By(MustBeValidUUID),
		),
	}.Filter()
}

func (r *getCommunityRequest) parse() *GetCommunityRequest {
	return &GetCommunityRequest{
		CommunityID: uuid.MustParse(r.CommunityID),
	}
}

func MustBeValidUUID(src interface{}) error {
	uuidRaw, ok := src.(string)
	if !ok {
		return errors.New("it is not a string")
	}

	_, err := uuid.Parse(uuidRaw)
	if err != nil {
		return errors.New("it is not a valid uuid")
	}

	return nil
}
