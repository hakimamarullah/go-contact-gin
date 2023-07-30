package repo

import (
	"context"
	"database/sql"

	"contact_ginv1/config"
	"contact_ginv1/domain/contract"
	"contact_ginv1/domain/model"
)

type GetPersonRepo struct {
	queryGet    string
	queryGetAll string
	dbs         *sql.DB
}

func NewGetPersonRepo(db *sql.DB) contract.GetPersonRepoInterface {
	return &GetPersonRepo{
		dbs:         db,
		queryGet:    "SELECT FirstName, LastName, Age, AddressId FROM Person WHERE Id = ?",
		queryGetAll: "SELECT FirstName, LastName, Age FROM Person",
	}
}

func (repo *GetPersonRepo) GetAllPerson() (person []model.Person, err error) {
	timeoutctx, cancel := context.WithTimeout(context.Background(), config.AppGetConfig().MysqlDB_TimeoutQuick)
	defer cancel()

	res, err := repo.dbs.QueryContext(timeoutctx, repo.queryGet)
	if err != nil {
		return
	}

	var temp model.Person
	for res.Next() {
		err = res.Scan(&temp.FirstName, &temp.LastName, &temp.Age)
		if err != nil {
			return nil, err
		}
		person = append(person, temp)
	}

	return
}

func (repo *GetPersonRepo) GetPersonById(id int) (person model.Person, err error) {
	timeoutctx, cancel := context.WithTimeout(context.Background(), config.AppGetConfig().MysqlDB_TimeoutQuick)
	defer cancel()

	res := repo.dbs.QueryRowContext(timeoutctx, repo.queryGet, id)
	err = res.Scan(&person.FirstName, &person.LastName, &person.Age, &person.AddressId)

	return
}
