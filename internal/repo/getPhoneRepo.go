package repo

import (
	"context"
	"database/sql"
	"time"

	"contact_ginv1/config"
	"contact_ginv1/domain/contract"
	"contact_ginv1/domain/model"
)

type GetPhoneRepo struct {
	queryGet          string
	queryGetAll       string
	queryGetByNumbers string
	queryGetByIMEI    string
	dbs               *sql.DB
}

func NewGetPhoneRepo(db *sql.DB) contract.GetPhoneRepoInterface {
	return &GetPhoneRepo{
		dbs:               db,
		queryGet:          "SELECT Numbers, IMEI FROM Phone WHERE Id = ?",
		queryGetAll:       "SELECT Numbers, IMEI FROM Phone",
		queryGetByNumbers: "SELECT Numbers, IMEI, PersonId FROM Phone WHERE UPPER(Numbers) = upper(?)",
		queryGetByIMEI:    "SELECT Numbers, IMEI, PersonId FROM Phone WHERE UPPER(IMEI) = upper(?)",
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

func (repo *GetPhoneRepo) GetPhoneByNumber(number string) (phone model.Phone, err error) {
	timeoutctx, cancel := context.WithTimeout(context.Background(), config.AppGetConfig().MysqlDB_TimeoutQuick)
	defer cancel()

	res := repo.dbs.QueryRowContext(timeoutctx, repo.queryGetByNumbers, number)
	err = res.Scan(&phone.Numbers, &phone.IMEI, &phone.PersonId)

	return
}

func (repo *GetPhoneRepo) GetPhoneByIMEI(imei string) (phone model.Phone, err error) {
	timeoutctx, cancel := context.WithTimeout(context.Background(), config.AppGetConfig().MysqlDB_TimeoutQuick)
	defer cancel()

	res := repo.dbs.QueryRowContext(timeoutctx, repo.queryGetByIMEI, imei)
	err = res.Scan(&phone.Numbers, &phone.IMEI, &phone.PersonId)

	return
}
