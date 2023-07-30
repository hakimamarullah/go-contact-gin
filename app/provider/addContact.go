package provider

import (
	"database/sql"

	"contact_ginv1/domain/contract"
	"contact_ginv1/internal/delivery/db"
	"contact_ginv1/internal/delivery/handler"
	"contact_ginv1/internal/repo"
	"contact_ginv1/internal/usecase"
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
	trxRepo := repo.NewTrxRepo(currentdb)

	usecase := usecase.NewAddContactUsecase(repoCountry, repoAddress, repoPerson, repoPhone, trxRepo)

	return handler.NewAddContactHandler(usecase)
}
