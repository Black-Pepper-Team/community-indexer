package postgres

import (
	"database/sql"

	"gitlab.com/distributed_lab/kit/pgdb"

	"github.com/black-pepper-team/community-indexer/internal/data"
)

type masterQ struct {
	db *pgdb.DB
}

func NewMasterQ(db *pgdb.DB) data.MasterQ {
	return &masterQ{
		db: db,
	}
}

func (q *masterQ) New() data.MasterQ {
	return NewMasterQ(q.db.Clone())
}

func (q *masterQ) CommunitiesQ() data.CommunitiesQ {
	return NewCommunitiesQ(q.db)
}

func (q *masterQ) Transaction(fn func() error) error {
	return q.db.Transaction(fn)
}

func (q *masterQ) IsolatedTransaction(isolationLevel sql.IsolationLevel, fn func() error) error {
	return q.db.TransactionWithOptions(&sql.TxOptions{Isolation: isolationLevel}, fn)
}
