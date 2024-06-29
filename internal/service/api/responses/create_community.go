package responses

import (
	"github.com/google/uuid"
)

type CreateCommunity struct {
	ID uuid.UUID `json:"id"`
}

func NewCreateCommunity(id uuid.UUID) *CreateCommunity {
	return &CreateCommunity{
		ID: id,
	}
}

func MockedCreateCommunity() *CreateCommunity {
	return &CreateCommunity{
		ID: uuid.MustParse("00000000-0000-0000-0000-000000000002"),
	}
}
