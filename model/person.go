package model

import (
	"github.com/mholt/binding"
)

type Person struct {
	ID         string            `gorethink:"id,omitempty"json:"id"`
	FirstName  string            `gorethink:"firstName"json:"firstName"`
	MiddleName string            `gorethink:"middleName"json:"middleName"`
	LastName   string            `gorethink:"lastName"json:"lastName"`
	Profiles   map[string]string `gorethink:",omitempty"json:",omitempty"`
	TimeStamp
}

func (p *Person) FieldMap() binding.FieldMap {
	return binding.FieldMap{
		&p.FirstName:  field("firstName", true),
		&p.MiddleName: field("middleName", false),
		&p.LastName:   field("lastName", true),
	}
}
