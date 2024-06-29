package responses

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"

	"github.com/black-pepper-team/community-indexer/internal/data"
	"github.com/black-pepper-team/community-indexer/internal/service/api/requests"
)

type Community struct {
	ID              uuid.UUID      `json:"id"`
	Name            string         `json:"name"`
	ContractAddress common.Address `json:"contract_address"`
	OwnerAddress    common.Address `json:"owner_address"`
}

func NewGetCommunity(community data.Community) *Community {
	return &Community{
		ID:              community.ID,
		Name:            community.Name,
		ContractAddress: community.ContractAddress,
		OwnerAddress:    community.OwnerAddress,
	}
}

func MockedGetCommunity(reg *requests.GetCommunityRequest) *Community {
	return &Community{
		ID:              reg.CommunityID,
		Name:            "Community 3",
		ContractAddress: common.HexToAddress("0x3"),
		OwnerAddress:    common.HexToAddress("0x300"),
	}
}
