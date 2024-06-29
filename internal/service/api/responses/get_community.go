package responses

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"

	"github.com/black-pepper-team/community-indexer/internal/data"
	"github.com/black-pepper-team/community-indexer/internal/service/api/requests"
)

type Community struct {
	ID              uuid.UUID      `json:"id"`
	Status          string         `json:"status"`
	Name            string         `json:"name"`
	Symbol          string         `json:"symbol"`
	ContractAddress common.Address `json:"contract_address"`
	OwnerAddress    common.Address `json:"owner_address"`
}

type GetCommunity struct {
	Community Community `json:"community"`
}

func NewGetCommunity(community *data.Community) *GetCommunity {
	return &GetCommunity{
		Community: Community{
			ID:              community.ID,
			Status:          string(community.Status),
			Symbol:          community.Symbol,
			Name:            community.Name,
			ContractAddress: community.ContractAddress,
			OwnerAddress:    community.OwnerAddress,
		},
	}
}

func MockedGetCommunity(reg *requests.GetCommunityRequest) *GetCommunity {
	return &GetCommunity{
		Community: Community{
			ID:              reg.CommunityID,
			Status:          "ready",
			Name:            "Community 3",
			Symbol:          "COM3",
			ContractAddress: common.HexToAddress("0x3"),
			OwnerAddress:    common.HexToAddress("0x300"),
		},
	}
}
