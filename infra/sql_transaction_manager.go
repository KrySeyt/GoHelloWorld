package infra

import (
	"database/sql"
)

type SqlTransactionManager struct {
	db *sql.DB
	tx *sql.Tx
}

func (self *SqlTransactionManager) Begin() {
	tx, err := self.db.Begin()
	if err != nil {
		panic(err)
	}

	self.tx = tx
}

func (self *SqlTransactionManager) Commit() {
	if self.tx == nil {
		panic("No started transaction for commit")
	}

	err := self.tx.Commit()
	if err != nil {
		panic(err)
	}

	self.tx = nil
}

func CreateSqlTransactionManager(db *sql.DB) *SqlTransactionManager {
	return &SqlTransactionManager{
		db: db,
	}
}
