package model

import (
	"net/http"

	"github.com/mholt/binding"
)

type Enrollment struct {
	ID        string `gorethink:"id,omitempty"json:"id"`
	StudentID string `gorethink:"studentId,omitempty"json:"studentId"`
	ClassID   string `gorethink:"classId,omitempty"json:"classId"`
	TermID    string `gorethink:"termId,omitempty"json:"termId"`
	TimeStamp
}

type EnrollmentAPIRes struct {
	ID        string  `gorethink:"id,omitempty"json:"id"`
	StudentID string  `gorethink:"studentId,omitempty"json:"studentId"`
	ClassID   string  `gorethink:"classId,omitempty"json:"classId"`
	TermID    string  `gorethink:"termId,omitempty"json:"termId"`
	Student   Student `gorethink:"student,omitempty"json:"student"`
	Person    Person  `gorethink:"person,omitempty"json:"person"`
	TimeStamp
}

func (e Enrollment) Validate(req *http.Request, errs binding.Errors) binding.Errors {
	if e.StudentID == "" {
		errs = append(errs, RequiredErr("studentId"))
	}
	if e.TermID == "" {
		errs = append(errs, RequiredErr("termId"))
	}
	if e.ClassID == "" {
		errs = append(errs, RequiredErr("classId"))
	}
	return errs
}

func (e *Enrollment) FieldMap() binding.FieldMap {
	return binding.FieldMap{
		&e.ID:        "id",
		&e.StudentID: "studentId",
		&e.TermID:    "termId",
		&e.ClassID:   "classId",
	}
}
