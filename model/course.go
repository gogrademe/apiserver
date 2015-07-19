package model

import (
	"net/http"

	"github.com/mholt/binding"
)

type Course struct {
	ID          string   `gorethink:"id,omitempty" json:"id"`
	Name        string   `gorethink:"name,omitempty" json:"name"`
	GradeLevel  string   `gorethink:"gradeLevel,omitempty" json:"gradeLevel"`
	MaxStudents int      `gorethink:"maxStudents,omitempty" json:"maxStudents"`
	Terms       []string `gorethink:"terms" json:"terms,omitempty"`
	TimeStamp
}

func (c Course) Validate(req *http.Request, errs binding.Errors) binding.Errors {
	if c.Name == "" {
		errs = append(errs, RequiredErr("name"))
	}
	if c.GradeLevel == "" {
		errs = append(errs, RequiredErr("gradeLevel"))
	}

	if len(c.Terms) <= 0 {
		errs = append(errs, RequiredErr("terms"))
	}
	return errs
}

func (c *Course) FieldMap(req *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&c.ID:         "id",
		&c.Name:       "name",
		&c.GradeLevel: "gradeLevel",
	}
}
