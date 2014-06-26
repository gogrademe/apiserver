package model

import (
	"github.com/mholt/binding"
)

type Teacher struct {
	ID          string `gorethink:"id,omitempty"json:"id"`
	PersonID    int64  `gorethink:"personId"json:"personId"`
	PhoneNumber string `gorethink:"phoneNumber"json:"personId"`
	Email       string `gorethink:"email"json:"email"`
	TimeStamp
}

func (t *Teacher) FieldMap() binding.FieldMap {
	return binding.FieldMap{
		&t.ID:       field("id", false),
		&t.PersonID: field("personId", true),
		&t.Email:    field("email", false),
	}
}
