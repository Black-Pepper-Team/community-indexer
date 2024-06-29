package data

import "database/sql"

type MasterQ interface {
	New() MasterQ

	CommunitiesQ() CommunitiesQ

	Transaction(func() error) error
	IsolatedTransaction(sql.IsolationLevel, func() error) error
}
