package db

import (
	"database/sql"
	"fmt"
	"time"

	"contact_ginv1/config"
	"contact_ginv1/domain/contract"

	_ "github.com/go-sql-driver/mysql"
)

var mysqldb MysqlDB

type MysqlDB struct {
	Dbs *sql.DB
	Trx *sql.Tx
}

func GetMysqlConnection() contract.Dbs {
	var nillDB MysqlDB
	if mysqldb != nillDB {
		return &mysqldb
	} else {
		mysqldb.OpenConnection()
		return &mysqldb
	}
}

func (d *MysqlDB) OpenConnection() {
	appConfig := config.AppGetConfig()

	username := appConfig.MysqlDB_USER
	password := appConfig.MysqlDB_PASS
	host := appConfig.MysqlDB_HOST
	port := appConfig.MysqlDB_PORT
	db := appConfig.MysqlDB_DB

	connetionstring := "%s:%s@tcp(%s:%s)/%s?parseTime=true"
	db_, err := sql.Open("mysql", fmt.Sprintf(connetionstring, username, password, host, port, db))
	if err != nil {
		panic(err)
	}

	db_.SetConnMaxIdleTime(10 * time.Minute)
	db_.SetConnMaxLifetime(12 * time.Hour)
	db_.SetMaxIdleConns(10)
	db_.SetMaxOpenConns(100)

	err = db_.Ping()
	if err != nil {
		panic(err)
	}

	d.Dbs = db_
}

func (d *MysqlDB) StartTrx() {
	trxs, err := mysqldb.Dbs.Begin()
	if err != nil {
		return
	}
	mysqldb.Trx = trxs
}

func (d *MysqlDB) DoneTrx(err error) {
	if err != nil {
		mysqldb.Trx.Rollback()
		mysqldb.Trx = &sql.Tx{}
	} else {
		mysqldb.Trx.Commit()
		mysqldb.Trx = &sql.Tx{}
	}
}
