package responses

import "github.com/black-pepper-team/community-indexer/internal/service/core"

func NewGetRegisterStatus(registerRequest *core.RegisterRequest) *RegisterRequest {
	return &RegisterRequest{
		Id:     registerRequest.Id,
		Status: string(registerRequest.Status),
	}
}
