package postgres

import (
	"database/sql"

	sq "github.com/Masterminds/squirrel"
	"github.com/ethereum/go-ethereum/common"
	"github.com/fatih/structs"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gitlab.com/distributed_lab/kit/pgdb"

	"github.com/black-pepper-team/community-indexer/internal/data"
)

const (
	committedStatesTableName = "communities"

	idColumnName              = "id"
	contractAddressColumnName = "contract_address"
	ownerAddressColumnName    = "owner_address"
)

type communitiesQ struct {
	db  *pgdb.DB
	sel sq.SelectBuilder
}

func NewCommunitiesQ(db *pgdb.DB) data.CommunitiesQ {
	return &communitiesQ{
		db:  db,
		sel: sq.Select("*").From(committedStatesTableName),
	}
}

func (q *communitiesQ) New() data.CommunitiesQ {
	return NewCommunitiesQ(q.db.Clone())
}

func (q *communitiesQ) Select() ([]data.Community, error) {
	var result []data.Community
	err := q.db.Select(&result, q.sel)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, errors.Wrap(err, "failed to select rows")
	}

	return result, nil
}

func (q *communitiesQ) Get() (*data.Community, error) {
	var result data.Community

	if err := q.db.Get(&result, q.sel); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, errors.Wrap(err, "failed to select rows")
	}

	return &result, nil
}

func (q *communitiesQ) Insert(community *data.Community) error {
	err := q.db.Exec(sq.Insert(committedStatesTableName).SetMap(structs.Map(community)))
	if err != nil {
		return errors.Wrap(err, "failed to insert rows")
	}

	return nil
}

func (q *communitiesQ) Update(community *data.Community) error {
	err := q.db.Exec(
		sq.Update(committedStatesTableName).
			SetMap(structs.Map(community)).
			Where(sq.Eq{idColumnName: community.ID}),
	)
	if err != nil {
		return errors.Wrap(err, "failed to update rows")
	}

	return nil
}

func (q *communitiesQ) WhereID(id ...uuid.UUID) data.CommunitiesQ {
	q.sel = q.sel.Where(sq.Eq{idColumnName: id})
	return q
}

func (q *communitiesQ) WhereContractAddress(address ...common.Address) data.CommunitiesQ {
	q.sel = q.sel.Where(sq.Eq{contractAddressColumnName: address})
	return q
}

func (q *communitiesQ) WhereOwnerAddress(address ...common.Address) data.CommunitiesQ {
	q.sel = q.sel.Where(sq.Eq{ownerAddressColumnName: address})
	return q
}
