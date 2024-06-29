package responses

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"

	"github.com/black-pepper-team/community-indexer/internal/data"
)

type CreateCommunity struct {
	Community Community `json:"community"`
}

func NewCreateCommunity(community *data.Community) *CreateCommunity {
	return &CreateCommunity{
		Community: Community{
			ID:              community.ID,
			Status:          string(community.Status),
			Name:            community.Name,
			ContractAddress: community.ContractAddress,
			OwnerAddress:    community.OwnerAddress,
		},
	}
}

func MockedCreateCommunity() *CreateCommunity {
	return &CreateCommunity{
		Community: Community{
			ID:              uuid.MustParse("00000000-0000-0000-0000-000000000003"),
			Status:          "ready",
			Name:            "Community 3",
			ContractAddress: common.HexToAddress("0x3"),
			OwnerAddress:    common.HexToAddress("0x300"),
		},
	}
}
