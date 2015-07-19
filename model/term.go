package model

import (
	"net/http"
	"time"

	"github.com/mholt/binding"
)

type Term struct {
	ID         string     `gorethink:"id,omitempty" json:"id"`
	Name       string     `gorethink:"name,omitempty" json:"name"`
	SchoolYear SchoolYear `gorethink:"schoolYear,omitempty" json:"schoolYear"`
	StartDate  time.Time  `gorethink:"startDate,omitempty" json:"startDate"`
	EndDate    time.Time  `gorethink:"endDate,omitempty" json:"endDate"`
	TimeStamp
}

func (t Term) Validate(req *http.Request, errs binding.Errors) binding.Errors {
	if t.Name == "" {
		errs = append(errs, RequiredErr("name"))
	}
	if t.StartDate.IsZero() {
		errs = append(errs, RequiredErr("startDate"))
	}
	if t.EndDate.IsZero() {
		errs = append(errs, RequiredErr("endDate"))
	}
	return errs
}

func (t *Term) FieldMap(req *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&t.ID:         "id",
		&t.Name:       "name",
		&t.SchoolYear: "schoolYear",
		&t.StartDate:  "startDate",
		&t.EndDate:    "endDate",
	}
}
