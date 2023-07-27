package repo

import (
	"context"
	"database/sql"
	"time"

	"contact_chiv2/domain/contract"
	"contact_chiv2/domain/model"
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
	timeoutctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := trx.ExecContext(timeoutctx, repo.query, data.FirstName, data.LastName, data.Age, data.AddressId)

	lastinserted, err = res.LastInsertId()

	return
}
