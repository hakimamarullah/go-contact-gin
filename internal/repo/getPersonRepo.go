package repo

import (
	"context"
	"database/sql"

	"contact_chiv2/config"
	"contact_chiv2/domain/contract"
	"contact_chiv2/domain/model"
)

type GetPersonRepo struct {
	queryGet    string
	queryGetAll string
	dbs         *sql.DB
}

func NewGetPersonRepo(db *sql.DB) contract.GetPersonRepoInterface {
	return &GetPersonRepo{
		dbs:         db,
		queryGet:    "",
		queryGetAll: "",
	}
}

func (repo *GetPersonRepo) GetAllPerson() (person []model.Person, err error) {
	timeoutctx, cancel := context.WithTimeout(context.Background(), config.AppGetConfig().MysqlDB_TimeoutQuick)
	defer cancel()

	res, err := repo.dbs.QueryContext(timeoutctx, repo.queryGetAll)
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
	err = res.Scan(&person.FirstName, &person.LastName, &person.Age)

	return
}
