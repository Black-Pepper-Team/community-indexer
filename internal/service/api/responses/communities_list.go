package responses

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"

	"github.com/black-pepper-team/community-indexer/internal/data"
)

type CommunitiesList struct {
	Communities []Community `json:"communities,omitempty"`
}

func NewCommunitiesList(communities []data.Community) *CommunitiesList {
	var rawCommunities []Community
	for _, community := range communities {
		rawCommunities = append(rawCommunities, Community{
			ID:              community.ID,
			Status:          string(community.Status),
			Name:            community.Name,
			Symbol:          community.Symbol,
			ContractAddress: community.ContractAddress,
			OwnerAddress:    community.OwnerAddress,
		})
	}

	return &CommunitiesList{
		Communities: rawCommunities,
	}
}

func MockedCommunitiesList() *CommunitiesList {
	return &CommunitiesList{
		Communities: []Community{
			{
				ID:              uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				Name:            "Community 1",
				Symbol:          "COM1",
				ContractAddress: common.HexToAddress("0x1"),
				OwnerAddress:    common.HexToAddress("0x100"),
				Status:          "ready",
			},
			{
				ID:              uuid.MustParse("00000000-0000-0000-0000-000000000002"),
				Name:            "Community 2",
				Symbol:          "COM2",
				ContractAddress: common.HexToAddress("0x2"),
				OwnerAddress:    common.HexToAddress("0x200"),
				Status:          "ready",
			},
			{
				ID:              uuid.MustParse("00000000-0000-0000-0000-000000000003"),
				Name:            "Community 3",
				Symbol:          "COM3",
				ContractAddress: common.HexToAddress("0x3"),
				OwnerAddress:    common.HexToAddress("0x300"),
				Status:          "deploying",
			},
			{
				ID:              uuid.MustParse("00000000-0000-0000-0000-000000000004"),
				Name:            "Community 4",
				Symbol:          "COM4",
				ContractAddress: common.HexToAddress("0x4"),
				OwnerAddress:    common.HexToAddress("0x400"),
				Status:          "deploy-failed",
			},
		},
	}
}
