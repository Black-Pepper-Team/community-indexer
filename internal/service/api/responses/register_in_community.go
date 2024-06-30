package responses

import (
	"github.com/google/uuid"

	"github.com/black-pepper-team/community-indexer/internal/service/core"
)

type RegisterRequest struct {
	Id     uuid.UUID `json:"id"`
	Status string    `json:"status"`
}

func NewRegisterInCommunity(registerRequest *core.RegisterRequest) *RegisterRequest {
	return &RegisterRequest{
		Id:     registerRequest.Id,
		Status: string(registerRequest.Status),
	}
}
