package repo

import (
	"context"
	"database/sql"

	"contact_ginv1/config"
	"contact_ginv1/domain/contract"
	"contact_ginv1/domain/model"
)

type GetAddressRepo struct {
	queryGet    string
	queryGetAll string
	dbs         *sql.DB
}

func NewGetAddressRepo(db *sql.DB) contract.GetAddressRepoInterface {
	return &GetAddressRepo{
		dbs:         db,
		queryGet:    "SELECT FullAddress, DistrictNumber, CountryId FROM Address WHERE Id = ?",
		queryGetAll: "SELECT FullAddress, DistrictNumber FROM Address",
	}
}

func (repo *GetAddressRepo) GetAllAddress() (address []model.Address, err error) {
	timeoutctx, cancel := context.WithTimeout(context.Background(), config.AppGetConfig().MysqlDB_TimeoutQuick)
	defer cancel()

	res, err := repo.dbs.QueryContext(timeoutctx, repo.queryGetAll)
	if err != nil {
		return
	}

	var temp model.Address
	for res.Next() {
		err = res.Scan(&temp.FullAddress, &temp.DistrictNumber)
		if err != nil {
			return nil, err
		}
		address = append(address, temp)
	}

	return
}

func (repo *GetAddressRepo) GetAddressById(id int) (address model.Address, err error) {
	timeoutctx, cancel := context.WithTimeout(context.Background(), config.AppGetConfig().MysqlDB_TimeoutQuick)
	defer cancel()

	res := repo.dbs.QueryRowContext(timeoutctx, repo.queryGet, id)
	err = res.Scan(&address.FullAddress, &address.DistrictNumber, &address.CountryId)

	return
}
