package model

import (
	// "errors"
	// "log"
	"github.com/mholt/binding"
	"net/http"
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
		&p.ID:         "id",
		&p.FirstName:  "firstName",
		&p.MiddleName: "middleName",
		&p.LastName:   "lastName",
	}
}

func (p Person) Validate(r *http.Request, errs binding.Errors) binding.Errors {
	if p.FirstName == "" {
		errs = append(errs, binding.Error{
			FieldNames: []string{"firstName"},
			Message:    "Required",
		})
	}
	if p.LastName == "" {
		errs = append(errs, binding.Error{
			FieldNames: []string{"lastName"},
			Message:    "Required",
		})
	}
	p.UpdateTime()
	return errs
}
