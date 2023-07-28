package repo

import (
	"context"
	"database/sql"

	"contact_chiv2/config"
	"contact_chiv2/domain/contract"
	"contact_chiv2/domain/model"
)

type GetAllContactRepo struct {
	query string
	dbs   *sql.DB
}

func NewGetAllContactRepo(db *sql.DB) contract.GetAllContactRepoInterface {
	return &GetAllContactRepo{
		dbs: db,
		query: `select p.Numbers , p.IMEI , p2.FirstName ,p2.LastName ,p2.Age ,a.FullAddress ,a.DistrictNumber ,c.CountryName ,c.ProvinceName  from Phone p 
		inner join Person p2 ON p.PersonId  = p2.Id
		inner  join Address a on p2.AddressId = a.Id 
		inner join Country c on a.CountryId = c.Id `,
	}
}

func (repo *GetAllContactRepo) GetAllContact() (contacts []model.Contact, err error) {
	timeoutctx, cancel := context.WithTimeout(context.Background(), config.AppGetConfig().MysqlDB_TimeoutQuick)
	defer cancel()

	res, err := repo.dbs.QueryContext(timeoutctx, repo.query)
	if err != nil {
		return
	}

	var temp model.Contact
	for res.Next() {
		err = res.Scan(&temp.Numbers, &temp.IMEI, &temp.FirstName, &temp.LastName, &temp.Age, &temp.FullAddress, &temp.DistrictNumber, &temp.CountryName, &temp.Region)
		if err != nil {
			return nil, err
		}

		contacts = append(contacts, temp)
	}

	return
}
