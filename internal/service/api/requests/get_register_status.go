package requests

import (
	"net/http"

	"github.com/go-chi/chi"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/uuid"
)

const (
	registerIdPathParam = "register-id"
)

type getRegisterRequest struct {
	RegisterID string
}

type GetRegisterRequest struct {
	RegisterID uuid.UUID
}

func NewGetRegister(r *http.Request) (*GetRegisterRequest, error) {
	requestRaw := getRegisterRequest{
		RegisterID: chi.URLParam(r, registerIdPathParam),
	}

	if err := requestRaw.validate(); err != nil {
		return nil, err
	}

	return requestRaw.parse(), nil
}

// nolint
func (r *getRegisterRequest) validate() error {
	return validation.Errors{
		"path/{register-id}": validation.Validate(
			r.RegisterID, validation.Required, validation.By(MustBeValidUUID),
		),
	}.Filter()
}

func (r *getRegisterRequest) parse() *GetRegisterRequest {
	return &GetRegisterRequest{
		RegisterID: uuid.MustParse(r.RegisterID),
	}
}
