package responses

import (
	"github.com/google/uuid"
)

type ImportCommunity struct {
	ID uuid.UUID `json:"id"`
}

func NewImportCommunity(id uuid.UUID) *CreateCommunity {
	return &CreateCommunity{
		ID: id,
	}
}

func MockedImportCommunity() *CreateCommunity {
	return &CreateCommunity{
		ID: uuid.MustParse("00000000-0000-0000-0000-000000000002"),
	}
}
