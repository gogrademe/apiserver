package model

import (
	"github.com/mholt/binding"
)

type Person struct {
	ID         string            `gorethink:"id,omitempty"json:"id"`
	FirstName  string            `gorethink:"firstName"json:"firstName"`
	MiddleName string            `gorethink:"middleName"json:"middleName"`
	LastName   string            `gorethink:"lastName"json:"lastName"`
	Profiles   map[string]string `gorethink:"profiles,omitempty"json:"profiles,omitempty"`
	TimeStamp
}

func (p *Person) FieldMap() binding.FieldMap {
	return binding.FieldMap{
		&p.ID:         field("id", false),
		&p.FirstName:  field("firstName", true),
		&p.MiddleName: field("middleName", false),
		&p.LastName:   field("lastName", true),
	}
}
