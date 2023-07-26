package db

import (
	"database/sql"
	"fmt"
	"time"

	"contact_chiv2/config"
	"contact_chiv2/domain/contract"
)

var postgredb PostgreDB

type PostgreDB struct {
	Dbs *sql.DB
}

func GetPostgreConnection() contract.Dbs {
	var nillDB PostgreDB
	if postgredb != nillDB {
		return &mysqldb
	} else {
		mysqldb.OpenConnection()
		return &mysqldb
	}
}

func (d *PostgreDB) OpenConnection() {
	username := config.MysqlDB_USER
	password := config.MysqlDB_PASS
	host := config.MysqlDB_HOST
	port := config.MysqlDB_PORT
	db := config.MysqlDB_DB

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
