package model

import (
// "errors"
// "log"

)

type Person struct {
	ID         string            `gorethink:"id,omitempty"json:"id"`
	FirstName  string            `gorethink:"firstName"json:"firstName"`
	MiddleName string            `gorethink:"middleName"json:"middleName"`
	LastName   string            `gorethink:"lastName"json:"lastName"`
	Profiles   map[string]string `gorethink:",omitempty"json:",omitempty"`
	TimeStamp
}

// func (p *Person) FieldMap() binding.FieldMap {
// 	return binding.FieldMap{
// 		&cf.User.ID: "user_id",
// 		&cf.Email:   "email",
// 		&cf.Message: binding.Field{
// 			Form:     "message",
// 			Required: true,
// 		},
// 	}
// }

func (t *Person) Validate() bool {
	if t.FirstName == "" {
		return false
	}
	if t.LastName == "" {
		return false
	}
	t.UpdateTime()
	return true
}
