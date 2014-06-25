package model

import (
	"github.com/mholt/binding"
)

type Teacher struct {
	PersonID    int64  `gorethink:"personId"json:"personId"`
	PhoneNumber string `gorethink:"phoneNumber"json:"personId"`
	Email       string `gorethink:"email"json:"email"`
	TimeStamp
}

func (t *Teacher) FieldMap() binding.FieldMap {
	return binding.FieldMap{
		&t.PersonID: field("personId", true),
		&t.Email:    field("email", false),
	}
}
