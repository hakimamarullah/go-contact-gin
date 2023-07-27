package repo

import (
	"context"
	"database/sql"
	"time"

	"contact_chiv2/config"
	"contact_chiv2/domain/contract"
	"contact_chiv2/domain/model"
)

type GetPhoneRepo struct {
	queryGet    string
	queryGetAll string
	dbs         *sql.DB
}

func NewGetPhoneRepo(db *sql.DB) contract.GetPhoneRepoInterface {
	return &GetPhoneRepo{
		dbs:         db,
		queryGet:    "",
		queryGetAll: "",
	}
}

func (repo *GetPhoneRepo) GetAllPhone() (person []model.Phone, err error) {
	timeoutctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := repo.dbs.QueryContext(timeoutctx, repo.queryGetAll)
	if err != nil {
		return
	}

	var temp model.Phone
	for res.Next() {
		err = res.Scan(&temp.Numbers, &temp.IMEI)
		if err != nil {
			return nil, err
		}
		person = append(person, temp)
	}

	return
}

func (repo *GetPhoneRepo) GetPhoneById(id int) (phone model.Phone, err error) {
	timeoutctx, cancel := context.WithTimeout(context.Background(), config.AppGetConfig().MysqlDB_TimeoutQuick)
	defer cancel()

	res := repo.dbs.QueryRowContext(timeoutctx, repo.queryGet, id)
	err = res.Scan(&phone.Numbers, &phone.IMEI)

	return
}
