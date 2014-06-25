package model

import (
// "errors"
// "log"
)

type Person struct {
	ID         string `gorethink:"id,omitempty"`
	FirstName  string `gorethink:"firstName"json:"firstName"`
	MiddleName string `gorethink:"middleName"json:"middleName"`
	LastName   string `gorethink:"lastName"json:"lastName"`
	Profiles   map[string]string
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
