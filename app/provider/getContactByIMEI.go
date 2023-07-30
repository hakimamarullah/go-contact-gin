package provider

import (
	"contact_ginv1/domain/contract"
	"contact_ginv1/internal/delivery/db"
	"contact_ginv1/internal/delivery/handler"
	"contact_ginv1/internal/repo"
	"contact_ginv1/internal/usecase"
	"database/sql"
)

func NewGetContactByIMEIHandler() contract.MainHandlerInterface {
	dbs := db.GetMysqlConnection()
	var currentdb *sql.DB

	switch v := dbs.(type) {
	case *db.MysqlDB:
		currentdb = v.Dbs
	case *db.PostgreDB:
		currentdb = v.Dbs
	}

	repoCountry := repo.NewGetCountryRepo(currentdb)
	repoAddress := repo.NewGetAddressRepo(currentdb)
	repoPerson := repo.NewGetPersonRepo(currentdb)
	repoPhone := repo.NewGetPhoneRepo(currentdb)

	getContactUsecase := usecase.NewGetContactUsecase(repoCountry, repoAddress, repoPerson, repoPhone)

	return handler.NewGetContactByIMEI(getContactUsecase)
}
