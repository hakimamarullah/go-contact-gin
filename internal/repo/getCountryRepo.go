package repo

import (
	"context"
	"database/sql"

	"contact_ginv1/config"
	"contact_ginv1/domain/contract"
	"contact_ginv1/domain/model"
)

type GetCountryRepo struct {
	queryGet    string
	queryGetAll string
	dbs         *sql.DB
}

func NewGetCountryRepo(db *sql.DB) contract.GetCountryRepoInterface {
	return &GetCountryRepo{
		dbs:         db,
		queryGet:    "SELECT CountryName, ProvinceName from Country WHERE Id = ?",
		queryGetAll: "SELECT CountryName, ProvinceName from Country",
	}
}

func (repo *GetCountryRepo) GetAllCountry() (country []model.Country, err error) {
	timeoutctx, cancel := context.WithTimeout(context.Background(), config.AppGetConfig().MysqlDB_TimeoutQuick)
	defer cancel()

	res, err := repo.dbs.QueryContext(timeoutctx, repo.queryGetAll)
	if err != nil {
		return
	}

	for res.Next() {
		temp := model.Country{}
		err := res.Scan(&temp.CountryName, &temp.Region)
		if err != nil {
			return nil, err
		}

		country = append(country, temp)
	}

	return
}

func (repo *GetCountryRepo) GetCountryById(id int) (country model.Country, err error) {
	timeoutctx, cancel := context.WithTimeout(context.Background(), config.AppGetConfig().MysqlDB_TimeoutQuick)
	defer cancel()

	res := repo.dbs.QueryRowContext(timeoutctx, repo.queryGet, id)
	err = res.Scan(&country.CountryName, &country.Region)

	return
}
