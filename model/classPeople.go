package model

import (
	"net/http"

	"github.com/mholt/binding"
)

type ClassPerson struct {
	ID       string   `gorethink:"id,omitempty"json:"id"`
	PersonID string   `gorethink:"personId,omitempty"json:"personId"`
	ClassID  string   `gorethink:"classId,omitempty"json:"classId"`
	TermID   string   `gorethink:"termId,omitempty"json:"termId"`
	Type     []string `gorethink:"type,omitempty"json:"type"`
	TimeStamp
}

func (c ClassPerson) Validate(req *http.Request, errs binding.Errors) binding.Errors {
	if c.PersonID == "" {
		errs = append(errs, RequiredErr("personId"))
	}
	if c.TermID == "" {
		errs = append(errs, RequiredErr("termId"))
	}
	if c.ClassID == "" {
		errs = append(errs, RequiredErr("classId"))
	}
	return errs
}

func (c *ClassPerson) FieldMap() binding.FieldMap {
	return binding.FieldMap{
		&c.ID:       "id",
		&c.PersonID: "personId",
		&c.TermID:   "termId",
		&c.ClassID:  "classId",
	}
}
