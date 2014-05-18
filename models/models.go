package models

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"time"
)

var (
	db *sqlx.DB
)

type AutoFields struct {
	Id        int
	CreatedAt time.Time `db:"createdAt"`
	UpdatedAt time.Time `db:"updatedAt"`
}

func (a *AutoFields) UpdateAuto() {

	t := time.Now().UTC()
	if a.Id != 0 {
		a.UpdatedAt = t
		return
	}
	a.CreatedAt = t
	a.UpdatedAt = t
	return
}
func init() {
	var err error
	// db, err = sqlx.Open("mysql", "root@tcp(127.0.0.1:3306)/dev_GoGrade?parseTime=true&loc=Local")
	db, err = sqlx.Open("postgres", "user=Matt dbname=dev_goGrade2 sslmode=disable")
	if err != nil {
		panic(fmt.Sprintf("Got error when connect database, the error is '%v'", err))
	}
	err = db.Ping()

	if err != nil {
		log.Fatalln(err)
	}
}

func SetupDB() error {

	_, err := CreateUser("test@test.com", "somePassword", "Admin")
	if err != nil {
		log.Println(err)
	}

	return nil
}

// /* Validation */
// func requiredString(field ...string) error {
// 	//Check if set

// 	// Then return err?
// }
