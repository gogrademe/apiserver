package model

import (
	"net/http"

	"github.com/mholt/binding"
)

// AssignmentType ...
type AssignmentType struct {
	ID         string   `gorethink:"id,omitempty"json:"id"`
	Name       string   `gorethink:"name,omitempty"json:"name"`
	Weight     float64  `gorethink:"weight,omitempty"json:"weight"`
	SubjectIDs []string `gorethink:"subjectIDs,omitempty"json:"subjectIDs,omitempty"`
	TimeStamp
}

// FieldMap ...
func (a *AssignmentType) FieldMap() binding.FieldMap {
	return binding.FieldMap{
		&a.ID:     "id",
		&a.Name:   "name",
		&a.Weight: "weight",
	}
}

// Validate ...
func (a AssignmentType) Validate(req *http.Request, errs binding.Errors) binding.Errors {
	if a.Name == "" {
		errs = append(errs, RequiredErr("name"))
	}

	if a.Weight > 1 || a.Weight < 0.005 {
		errs = append(errs, binding.Error{
			FieldNames: []string{"weight"},
			Message:    "must be between .5% and 100%",
		})
	}
	return errs
}
