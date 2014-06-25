package model

import (
// "errors"
// "log"
)

type Person struct {
	ID         int64  `json:"id"`
	FirstName  string `db:"first_name"json:"firstName"`
	MiddleName string `db:"middle_name"json:"middleName"`
	LastName   string `db:"last_name"json:"lastName"`
	TimeStamp
}

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
