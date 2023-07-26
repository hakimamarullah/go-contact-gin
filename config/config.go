package config

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	MysqlDB_USER string = ""
	MysqlDB_PASS string = ""
	MysqlDB_HOST string = ""
	MysqlDB_PORT string = ""
	MysqlDB_DB   string = ""
)

func init() {
	godotenv.Load("/home/t1ramisu/Project/Bootcamp/batch3/week4/contact_chi/project.env")
	MysqlDB_USER = os.Getenv("user")
	MysqlDB_PASS = os.Getenv("password")
	MysqlDB_HOST = os.Getenv("host")
	MysqlDB_PORT = os.Getenv("port")
	MysqlDB_DB = os.Getenv("db")
}
