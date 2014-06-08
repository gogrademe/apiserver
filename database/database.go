package database

import (
	"github.com/jmoiron/sqlx"
	// Imports pq sql implimentation.
	_ "github.com/lib/pq"
	"log"
)

var (
	db *sqlx.DB
)

// Init connects to the db.
func Init(name, datasource string) error {
	var err error

	db, err = sqlx.Open(name, datasource)
	if err != nil {
		return err
	}
	err = db.Ping()

	if err != nil {
		return err
	}
	return nil
}

// SetupDB will be used to bootstrap the DB
func SetupDB() error {

	_, err := CreateUser("test@test.com", "somePassword", "Admin")
	if err != nil {
		log.Println(err)
	}

	return nil
}
