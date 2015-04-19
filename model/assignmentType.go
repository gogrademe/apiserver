package model

import (
	"net/http"

	"github.com/mholt/binding"
)

// AssignmentType ...
type AssignmentType struct {
	ID      string  `gorethink:"id,omitempty" json:"id"`
	Name    string  `gorethink:"name" json:"name"`
	Weight  float64 `gorethink:"weight" json:"weight"`
	ClassID string  `gorethink:"classId" json:"classID"`
	TimeStamp
}

// FieldMap ...
func (a *AssignmentType) FieldMap() binding.FieldMap {
	return binding.FieldMap{
		&a.ID:     "id",
		&a.Name:   "name",
		&a.Name:   "weight",
		&a.Weight: "weight",
	}
}

// Validate ...
func (a AssignmentType) Validate(req *http.Request, errs binding.Errors) binding.Errors {
	if a.Name == "" {
		errs = append(errs, RequiredErr("name"))
	}

	if a.ClassID == "" {
		errs = append(errs, RequiredErr("classId"))
	}

	if a.Weight > 1 || a.Weight < 0.005 {
		errs = append(errs, binding.Error{
			FieldNames: []string{"weight"},
			Message:    "must be between .5% and 100%",
		})
	}
	return errs
}
