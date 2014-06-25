package model

type Teacher struct {
	PersonID    int64  `gorethink:"personId"json:"personId"`
	PhoneNumber string `gorethink:"phoneNumber"json:"personId"`
	Email       string `gorethink:"email"json:"email"`
	TimeStamp
}
