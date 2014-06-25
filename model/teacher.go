package model

type Teacher struct {
	PersonID    int64  `gorethink:"personID"`
	PhoneNumber string `gorethink:"phoneNumber"`
	Email       string `gorethink:"email"`
	TimeStamp
}
