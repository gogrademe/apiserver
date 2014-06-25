package model

type TeacherProfile struct {
	PersonID    int64  `db:"person_id"`
	PhoneNumber string `db:"phone_number"`
	Email       string `db:"email"`
	TimeStamp
}
