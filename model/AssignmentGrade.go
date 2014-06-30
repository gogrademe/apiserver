package model

import (
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

// Validate ...
func (a *AssignmentGrade) Validate() *ValErrors {
	var v *ValErrors

	v.RequiredString(a.AssignmentID, "assignmentId")
	v.RequiredString(a.StudentID, "studentId")
	v.RequiredString(a.Grade, "grade")

	return v
}
