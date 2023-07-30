package provider

import (
	"database/sql"

	"contact_ginv1/domain/contract"
	"contact_ginv1/internal/delivery/db"
	"contact_ginv1/internal/delivery/handler"
	"contact_ginv1/internal/repo"
	"contact_ginv1/internal/usecase"
)

func NewGetAllHandler() contract.MainHandlerInterface {
	dbs := db.GetMysqlConnection()
	var currentdb *sql.DB

	switch v := dbs.(type) {
	case *db.MysqlDB:
		currentdb = v.Dbs
	case *db.PostgreDB:
		currentdb = v.Dbs
	}

	getAllContactRepo := repo.NewGetAllContactRepo(currentdb)
	usecase := usecase.NewGetAllContactUsecase(getAllContactRepo)

	return handler.NewGetAllContactHandler(usecase)
}
