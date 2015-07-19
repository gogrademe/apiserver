package model

import (
	"net/http"

	"github.com/mholt/binding"
)

// AssignmentGroup ...
type AssignmentGroup struct {
	ID      string  `gorethink:"id,omitempty" json:"id"`
	Name    string  `gorethink:"name" json:"name"`
	Weight  float64 `gorethink:"weight" json:"weight"`
	ClassID string  `gorethink:"classId" json:"classId"`
	TermID  string  `gorethink:"termId" json:"termId"`
	TimeStamp
}

// FieldMap ...
func (a *AssignmentGroup) FieldMap(req *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&a.ID:      "id",
		&a.Name:    "name",
		&a.Weight:  "weight",
		&a.ClassID: "classId",
		&a.TermID:  "termId",
	}
}

// Validate ...
func (a AssignmentGroup) Validate(req *http.Request, errs binding.Errors) binding.Errors {
	if a.Name == "" {
		errs = append(errs, RequiredErr("name"))
	}

	if a.ClassID == "" {
		errs = append(errs, RequiredErr("classId"))
	}
	if a.TermID == "" {
		errs = append(errs, RequiredErr("termId"))
	}

	if a.Weight > 1 || a.Weight < 0.005 {
		errs = append(errs, binding.Error{
			FieldNames: []string{"weight"},
			Message:    "must be between .5% and 100%",
		})
	}
	return errs
}
