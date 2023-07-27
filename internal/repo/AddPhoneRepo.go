package repo

import (
	"context"
	"database/sql"
	"time"

	"contact_chiv2/domain/contract"
	"contact_chiv2/domain/model"
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
	timeoutctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := trx.ExecContext(timeoutctx, repo.query, data.PersonId, data.Numbers, data.IMEI)

	lastinserted, err = res.LastInsertId()

	return
}
