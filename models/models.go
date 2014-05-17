package models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var (
	db *sqlx.DB
)

func init() {
	var err error
	db, err = sqlx.Open("mysql", "root@tcp(127.0.0.1:3306)/dev_GoGrade?parseTime=true&loc=Local")
	if err != nil {
		panic(fmt.Sprintf("Got error when connect database, the error is '%v'", err))
	}
	err = db.Ping()

	if err != nil {
		panic(err)
	}
}

func SetupDB() error {

	_, err := CreateUser("test@test.com", "somePassword", "Admin")
	if err != nil {
		panic(err)
	}

	return nil
}
