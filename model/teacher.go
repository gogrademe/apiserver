package model

import (
	"github.com/mholt/binding"
)

type Teacher struct {
	ID          string `gorethink:"id,omitempty"json:"id"`
	PersonID    string `gorethink:"personId,omitempty"json:"personId"`
	PhoneNumber string `gorethink:"phoneNumber,omitempty"json:"personId"`
	Email       string `gorethink:"email,omitempty"json:"email"`
	TimeStamp
}

func (s *Teacher) Validate() *ValErrors {
	var v *ValErrors

	v.RequiredString(s.PersonID, "personId")

	return v
}
func (t *Teacher) FieldMap() binding.FieldMap {
	return binding.FieldMap{
		&t.ID:          field("id", false),
		&t.PersonID:    field("personId", true),
		&t.PhoneNumber: field("phoneNumber", false),
		&t.Email:       field("email", false),
	}
}
