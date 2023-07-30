package repo

import (
	"context"
	"database/sql"

	"contact_ginv1/config"
	"contact_ginv1/domain/contract"
	"contact_ginv1/domain/model"
)

type AddPhoneRepo struct {
	query string
	dbs   *sql.DB
}

func NewAddPhoneRepo(db *sql.DB) contract.AddPhoneRepoInterface {
	return &AddPhoneRepo{
		query: "INSERT INTO Phone(PersonId, Numbers, IMEI) VALUES(?, nullif(?,''), nullif(?,''))",
		dbs:   db,
	}
}

func (repo *AddPhoneRepo) AddPhone(data model.Phone) (lastinserted int64, err error) {
	timeoutctx, cancel := context.WithTimeout(context.Background(), config.AppGetConfig().MysqlDB_TimeoutQuick)
	defer cancel()

	res, err := trx.ExecContext(timeoutctx, repo.query, data.PersonId, data.Numbers, data.IMEI)

	lastinserted, err = res.LastInsertId()

	return
}
