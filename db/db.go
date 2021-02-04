package db

import (
	"database/sql"
	"ewarung-api-experiment/config"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func Init() {
	conf, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	connectionString := conf.DBUsername + ":" + conf.DBPassword + "@tcp(" + conf.DBHost + ":" + conf.DBPort + ")/" + conf.DBDatabase

	db, err = sql.Open("mysql", connectionString)
	if err != nil {
		panic("connectionString error..")
	}

	err = db.Ping()
	if err != nil {
		panic("DSN Invalid")
	}
}

func CreateCon() *sql.DB {
	return db
}
