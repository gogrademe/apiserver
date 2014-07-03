package model

import (
	"net/http"

	"github.com/mholt/binding"
)

type (
	Class struct {
		ID          string   `gorethink:"id,omitempty"json:"id"`
		Name        string   `gorethink:"name,omitempty"json:"name"`
		GradeLevel  string   `gorethink:"gradeLevel,omitempty"json:"gradeLevel"`
		Subject     string   `gorethink:"subject,omitempty"json:"subject"`
		MaxStudents int      `gorethink:"maxStudents,omitempty"json:"maxStudents"`
		Terms       []string `gorethink:"terms"json:"terms,omitempty"`
		TimeStamp
	}
)

func (c Class) Validate(req *http.Request, errs binding.Errors) binding.Errors {
	if c.Name == "" {
		errs = append(errs, RequiredErr("name"))
	}
	return errs
}

func (c *Class) FieldMap() binding.FieldMap {
	return binding.FieldMap{
		&c.ID:         "id",
		&c.Name:       "name",
		&c.GradeLevel: "gradeLevel",
		&c.Subject:    "subject",
	}
}
