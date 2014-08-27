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
	StudentID    string `gorethink:"studentId,omitempty"json:"studentId"`
	Grade        string `gorethink:"grade,omitempty"json:"grade"`
	TimeStamp
}

// FieldMap ...
func (a *AssignmentGrade) FieldMap() binding.FieldMap {
	return binding.FieldMap{
		&a.ID:           field("id", false),
		&a.AssignmentID: field("assignmentId", true),
		&a.StudentID:    field("studentId", true),
		&a.Grade:        field("grade", true),
	}
}
func (a AssignmentGrade) Validate(req *http.Request, errs binding.Errors) binding.Errors {
	if strings.TrimSpace(a.AssignmentID) == "" {
		errs = append(errs, RequiredErr("assignmentId"))
	}
	if strings.TrimSpace(a.StudentID) == "" {
		errs = append(errs, RequiredErr("studentId"))
	}
	if strings.TrimSpace(a.Grade) == "" {
		errs = append(errs, RequiredErr("grade"))
	}
	return errs
}
