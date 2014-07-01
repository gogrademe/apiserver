package model

import (
	"net/http"
	"time"

	"github.com/mholt/binding"
)

type (
	Class struct {
		ID         string            `gorethink:"id,omitempty"json:"id"`
		Name       string            `gorethink:"name,omitempty"json:"name"`
		GradeLevel string            `gorethink:"gradeLevel,omitempty"json:"gradeLevel"`
		Subject    string            `gorethink:"subject,omitempty"json:"subject"`
		Terms      map[string]string `gorethink:"-"json:"terms,omitempty"`
		TimeStamp
	}

	ClassTerm struct {
		ID          string       `gorethink:"id,omitempty"json:"id"`
		ClassID     string       `gorethink:"classId,omitempty"json:"classId"`
		Name        string       `gorethink:"name,omitempty"json:"name"`
		StartDate   time.Time    `gorethink:"startDate,omitempty"json:"startDate"`
		EndDate     time.Time    `gorethink:"endDate,omitempty"json:"endDate"`
		Assignments []Assignment `gorethink:"-"json:"assignments"`
		People      []Person     `gorethink:"-"json:"people"`
		TimeStamp
	}
)

func (c Class) Validate(req *http.Request, errs binding.Errors) binding.Errors {
	if c.Name == "" {
		errs = append(errs, RequiredErr("name"))
	}
	if c.Subject == "" {
		errs = append(errs, RequiredErr("subject"))
	}
	return errs
}
func (c ClassTerm) Validate(req *http.Request, errs binding.Errors) binding.Errors {
	if c.ClassID == "" {
		errs = append(errs, RequiredErr("classId"))
	}
	if c.Name == "" {
		errs = append(errs, RequiredErr("name"))
	}
	return errs
}

// func (c *Class) Validate() *ValErrors {
// 	var v *ValErrors
// 	v.RequiredString(c.Name, "name")
// 	v.RequiredString(c.Subject, "subject")
//
// 	return v
// }
// func (c *ClassTerm) Validate() *ValErrors {
// 	var v *ValErrors
// 	v.RequiredString(c.ClassID, "classId")
// 	return v
// }

func (c *Class) FieldMap() binding.FieldMap {
	return binding.FieldMap{
		&c.ID:         field("id", false),
		&c.Name:       field("name", true),
		&c.GradeLevel: field("gradeLevel", false),
		&c.Subject:    field("subject", true),
	}
}

func (ct *ClassTerm) FieldMap() binding.FieldMap {
	return binding.FieldMap{
		&ct.ID:        field("id", false),
		&ct.ClassID:   field("classId", true),
		&ct.Name:      field("name", true),
		&ct.StartDate: field("startDate", false),
		&ct.EndDate:   field("endDate", false),
	}
}
