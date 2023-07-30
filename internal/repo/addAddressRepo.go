package repo

import (
	"context"
	"database/sql"

	"contact_ginv1/config"
	"contact_ginv1/domain/contract"
	"contact_ginv1/domain/model"
)

type AddAddressRepo struct {
	query string
	dbs   *sql.DB
}

func NewAddresRepo(db *sql.DB) contract.AddAddressRepoInterface {
	return &AddAddressRepo{
		query: "INSERT INTO Address(CountryId, FullAddress, DistrictNumber) VALUES(?, nullif(?,''), nullif(?,0))",
		dbs:   db,
	}
}

func (repo *AddAddressRepo) AddAddress(data model.Address) (lastinserted int64, err error) {
	timeoutctx, cancel := context.WithTimeout(context.Background(), config.AppGetConfig().MysqlDB_TimeoutQuick)
	defer cancel()

	res, err := trx.ExecContext(timeoutctx, repo.query, data.CountryId, data.FullAddress, data.DistrictNumber)
	if err != nil {
		return
	}

	lastinserted, err = res.LastInsertId()
	return
}
