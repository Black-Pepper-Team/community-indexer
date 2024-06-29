package data

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
)

type CommunitiesQ interface {
	New() CommunitiesQ

	Get() (*Community, error)
	Select() ([]Community, error)
	Insert(*Community) error
	Update(*Community) error

	WhereID(...uuid.UUID) CommunitiesQ
	WhereContractAddress(...common.Address) CommunitiesQ
	WhereOwnerAddress(...common.Address) CommunitiesQ
}

type Community struct {
	ID              uuid.UUID       `db:"id"               structs:"id"`
	Status          CommunityStatus `db:"status"           structs:"status"`
	Name            string          `db:"name"             structs:"name"`
	Symbol          string          `db:"symbol"           structs:"symbol"`
	ContractAddress common.Address  `db:"contract_address" structs:"contract_address"`
	OwnerAddress    common.Address  `db:"owner_address"    structs:"owner_address"`
}

type CommunityStatus string

const (
	Ready          CommunityStatus = "ready"
	Deploying      CommunityStatus = "deploying"
	FailedToDeploy CommunityStatus = "deploy-failed"
)
