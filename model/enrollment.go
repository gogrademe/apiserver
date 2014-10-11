package model

import (
	"net/http"

	"github.com/mholt/binding"
)

type Enrollment struct {
	ID       string `gorethink:"id,omitempty"json:"id"`
	PersonID string `gorethink:"personId,omitempty"json:"personId"`
	ClassID  string `gorethink:"classId,omitempty"json:"classId"`
	TermID   string `gorethink:"termId,omitempty"json:"termId"`
	TimeStamp
}

type EnrollmentAPIRes struct {
	ID       string `gorethink:"id,omitempty"json:"id"`
	PersonID string `gorethink:"personId,omitempty"json:"personId"`
	ClassID  string `gorethink:"classId,omitempty"json:"classId"`
	TermID   string `gorethink:"termId,omitempty"json:"termId"`
	Person   Person `gorethink:"person,omitempty"json:"person"`
	TimeStamp
}

func (e Enrollment) Validate(req *http.Request, errs binding.Errors) binding.Errors {
	if e.PersonID == "" {
		errs = append(errs, RequiredErr("personId"))
	}
	if e.TermID == "" {
		errs = append(errs, RequiredErr("termId"))
	}
	if e.ClassID == "" {
		errs = append(errs, RequiredErr("classId"))
	}
	return errs
}

func (e *Enrollment) FieldMap() binding.FieldMap {
	return binding.FieldMap{
		&e.ID:       "id",
		&e.PersonID: "personId",
		&e.TermID:   "termId",
		&e.ClassID:  "classId",
	}
}
