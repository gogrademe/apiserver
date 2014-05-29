package database

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

var (
	db *sqlx.DB
)

func init() {
	var err error

	db, err = sqlx.Open("postgres", "user=Matt dbname=dev_go_grade sslmode=disable")
	if err != nil {
		log.Fatalln("Postgres:", err)
	}
	err = db.Ping()

	if err != nil {
		log.Fatalln("Postgres:", err)
	}
}

func SetupDB() error {

	_, err := CreateUser("test@test.com", "somePassword", "Admin")
	if err != nil {
		log.Println(err)
	}

	return nil
}
