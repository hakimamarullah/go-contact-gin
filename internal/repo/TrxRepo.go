package repo

import (
	"database/sql"

	"contact_ginv1/domain/contract"
	"contact_ginv1/internal/delivery/db"
)

var trx *sql.Tx

type Transaction struct {
	dbs *sql.DB
}

func NewTrxRepo(db *sql.DB) contract.TrxRepoInterface {
	return &Transaction{
		dbs: db,
	}
}

func (repo *Transaction) StartTrx() {
	dbs := db.GetMysqlConnection()
	dbs.StartTrx()

	currentdb, ok := dbs.(*db.MysqlDB)
	if !ok {
		return
	}

	trx = currentdb.Trx
}

func (repo *Transaction) DoneTrx(err error) {
	dbs := db.GetMysqlConnection()
	dbs.DoneTrx(err)

	currentdb, ok := dbs.(*db.MysqlDB)
	if !ok {
		return
	}

	trx = currentdb.Trx
}
