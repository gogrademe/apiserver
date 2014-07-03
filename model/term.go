package model

import (
	"net/http"
	"time"

	"github.com/mholt/binding"
)

type Term struct {
	ID        string    `gorethink:"id,omitempty"json:"id"`
	Name      string    `gorethink:"name,omitempty"json:"name"`
	StartDate time.Time `gorethink:"startDate,omitempty"json:"startDate"`
	EndDate   time.Time `gorethink:"endDate,omitempty"json:"endDate"`
	// Assignments []Assignment `gorethink:"-"json:"assignments"`
	// People      []Person     `gorethink:"-"json:"people"`
	TimeStamp
}

func (t Term) Validate(req *http.Request, errs binding.Errors) binding.Errors {
	if t.Name == "" {
		errs = append(errs, RequiredErr("name"))
	}
	return errs
}

func (t *Term) FieldMap() binding.FieldMap {
	return binding.FieldMap{
		&t.ID:   "id",
		&t.Name: "name",
	}
}
