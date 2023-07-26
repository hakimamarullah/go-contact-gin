package provider

import (
	"database/sql"

	"contact_chiv2/domain/contract"
	"contact_chiv2/internal/delivery/db"
	"contact_chiv2/internal/delivery/handler"
	"contact_chiv2/internal/repo"
	"contact_chiv2/internal/usecase"
)

func NewAddContactHandler() contract.MainHandlerInterface {
	dbs := db.GetMysqlConnection()
	var currentdb *sql.DB

	switch v := dbs.(type) {
	case *db.MysqlDB:
		currentdb = v.Dbs
	case *db.PostgreDB:
		currentdb = v.Dbs
	}

	repoCountry := repo.NewAddCountryRepo(currentdb)
	repoAddress := repo.NewAddresRepo(currentdb)
	repoPerson := repo.NewAddPersonRepo(currentdb)
	repoPhone := repo.NewAddPhoneRepo(currentdb)

	usecase := usecase.NewAddContactUsecase(repoCountry, repoAddress, repoPerson, repoPhone)

	return handler.NewAddContactHandler(usecase)
}
