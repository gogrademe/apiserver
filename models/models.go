package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var (
	db gorm.DB
)

func init() {
	var err error
	db, err = gorm.Open("postgres", "user=Matt dbname=dev_goGradeGorm1 sslmode=disable")
	if err != nil {
		panic(fmt.Sprintf("Got error when connect database, the error is '%v'", err))
	}

	// With it you could use package `database/sql`'s builtin methods
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.DB().Ping()
	db.SingularTable(true)

}

func SetupDB() error {

	db.AutoMigrate(User{})
	db.AutoMigrate(Class{})
	_, err := CreateUser("test@test.com", "somePassword", "Admin")
	if err != nil {
		panic(err)
	}

	return nil
}
