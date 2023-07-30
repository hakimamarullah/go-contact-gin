package repo

import (
	"context"
	"database/sql"

	"contact_ginv1/config"
	"contact_ginv1/domain/contract"
	"contact_ginv1/domain/model"
)

type AddPersonRepo struct {
	query string
	dbs   *sql.DB
}

func NewAddPersonRepo(db *sql.DB) contract.AddPersonRepoInterface {
	return &AddPersonRepo{
		query: "INSERT INTO Person(FirstName, LastName, Age, AddressId) VALUES(nullif(?,''), nullif(?,''), nullif(?,0), ?)",
		dbs:   db,
	}
}

func (repo *AddPersonRepo) AddPerson(data model.Person) (lastinserted int64, err error) {
	timeoutctx, cancel := context.WithTimeout(context.Background(), config.AppGetConfig().MysqlDB_TimeoutQuick)
	defer cancel()

	res, err := trx.ExecContext(timeoutctx, repo.query, data.FirstName, data.LastName, data.Age, data.AddressId)

	lastinserted, err = res.LastInsertId()

	return
}
