package services

import (
	"database/sql"
	"fmt"
)

var db *sql.DB
var err error

const (
	Host     = "localhost"
	Port     = 5432
	User     = "postgres"
	Password = "753159"
	Dbname   = "api-rest-snet"
)

func checkErr(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}

func Connection() {
	connection := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", Host, Port, User, Password, Dbname)

	db, err = sql.Open("postgres", connection)
	checkErr(err)
	err = db.Ping()
	checkErr(err)
}
