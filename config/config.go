package config

import (
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type configStruct struct {
	MysqlDB_USER         string
	MysqlDB_PASS         string
	MysqlDB_HOST         string
	MysqlDB_PORT         string
	MysqlDB_DB           string
	MysqlDB_TimeoutQuick time.Duration
	MysqlDB_TimeoutMid   time.Duration
	MysqlDB_TimeoutSlow  time.Duration
}

var appConfig = new(configStruct)

func init() {
	godotenv.Load("d:/phincon/go-contact-gin/project.env")
	appConfig.MysqlDB_USER = os.Getenv("user")
	appConfig.MysqlDB_PASS = os.Getenv("password")
	appConfig.MysqlDB_HOST = os.Getenv("host")
	appConfig.MysqlDB_PORT = os.Getenv("port")
	appConfig.MysqlDB_DB = os.Getenv("db")

	durationQuick, err := strconv.Atoi(os.Getenv("mysqltimeout1"))
	if err != nil {
		panic(err)
	}
	appConfig.MysqlDB_TimeoutQuick = time.Duration(durationQuick) * time.Second

	durationMid, err := strconv.Atoi(os.Getenv("mysqltimeout2"))
	if err != nil {
		panic(err)
	}
	appConfig.MysqlDB_TimeoutSlow = time.Duration(durationMid) * time.Second

	durationSlow, err := strconv.Atoi(os.Getenv("mysqltimeout3"))
	if err != nil {
		panic(err)
	}
	appConfig.MysqlDB_TimeoutSlow = time.Duration(durationSlow) * time.Second
}

func AppGetConfig() configStruct {
	return *appConfig
}
