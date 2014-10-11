package model

import (
	"net/http"
	"strings"

	"github.com/mholt/binding"
)

//AssignmentGrade ...
type AssignmentGrade struct {
	ID           string `gorethink:"id,omitempty"json:"id"`
	AssignmentID string `gorethink:"assignmentId,omitempty"json:"assignmentId"`
	PersonID     string `gorethink:"personId,omitempty"json:"personId"`
	Grade        string `gorethink:"grade,omitempty"json:"grade"`
	TimeStamp
}

// FieldMap ...
func (a *AssignmentGrade) FieldMap() binding.FieldMap {
	return binding.FieldMap{
		&a.ID:           "id",
		&a.AssignmentID: "assignmentId",
		&a.PersonID:     "personId",
		&a.Grade:        "grade",
	}
}
func (a AssignmentGrade) Validate(req *http.Request, errs binding.Errors) binding.Errors {
	if strings.TrimSpace(a.AssignmentID) == "" {
		errs = append(errs, RequiredErr("assignmentId"))
	}
	if strings.TrimSpace(a.PersonID) == "" {
		errs = append(errs, RequiredErr("personId"))
	}
	if strings.TrimSpace(a.Grade) == "" {
		errs = append(errs, RequiredErr("grade"))
	}
	return errs
}
