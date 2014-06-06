package model

import (
	"time"
)

// var (
// 	db *sqlx.DB
// )

// // type AutoFields struct {
// // 	Id        int64
// // 	CreatedAt time.Time `db:"created_at"`
// // 	UpdatedAt time.Time `db:"updated_at"`
// // }
type TimeStamp struct {
	CreatedAt time.Time `db:"created_at"json:"createdAt"`
	UpdatedAt time.Time `db:"updated_at"json:"updatedAt"`
}

func (a *TimeStamp) UpdateTime() {

	t := time.Now().UTC()
	if !a.CreatedAt.IsZero() {
		a.UpdatedAt = t
		return
	}
	a.CreatedAt = t
	a.UpdatedAt = t
	return
}

// func init() {
// 	var err error
// 	// db, err = sqlx.Open("mysql", "root@tcp(127.0.0.1:3306)/dev_GoGrade?parseTime=true&loc=Local")
// 	db, err = sqlx.Open("postgres", "user=Matt dbname=dev_go_grade sslmode=disable")
// 	if err != nil {
// 		log.Fatalln("Postgres:", err)
// 	}
// 	err = db.Ping()

// 	if err != nil {
// 		log.Fatalln("Postgres:", err)
// 	}
// }

// func SetupDB() error {

// 	_, err := CreateUser("test@test.com", "somePassword", "Admin")
// 	if err != nil {
// 		log.Println(err)
// 	}

// 	return nil
// }

// // /* Validation */
// // func requiredString(field ...string) error {
// // 	//Check if set

// // 	// Then return err?
// // }
