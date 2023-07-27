package repo

import (
	"database/sql"

	"contact_chiv2/domain/contract"
	"contact_chiv2/domain/model"
)

type GetCountryRepo struct {
	dbs *sql.DB
}

func NewGetCountryRepo(db *sql.DB) contract.GetCountryRepoInterface {
	return &GetCountryRepo{
		dbs: db,
	}
}

func (repo *GetCountryRepo) GetAllCountry() (country []model.Country, err error) {
	query := "SELECT CountryName, ProvinceName from Country"
	res, err := repo.dbs.Query(query)
	if err != nil {
		return
	}

	for res.Next() {
		temp := model.Country{}
		res.Scan(&temp.CountryName, &temp.Region)

		country = append(country, temp)
	}

	return
}

func (repo *AddPhoneRepo) DoneTrx(err error) {
	if err != nil {
		trx.Rollback()
		trx = &sql.Tx{}
	} else {
		trx.Commit()
		trx = &sql.Tx{}
	}
}
