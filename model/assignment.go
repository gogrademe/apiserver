package model

import (
	"net/http"
	"time"

	"github.com/mholt/binding"
)

//Assignment ...
type Assignment struct {
	ID       string         `gorethink:"id,omitempty" json:"id"`
	Name     string         `gorethink:"name,omitempty" json:"name"`
	ClassID  string         `gorethink:"classId,omitempty" json:"classId"`
	TermID   string         `gorethink:"termId,omitempty" json:"termId"`
	TypeID   string         `gorethink:"typeId,omitempty" json:"typeId"`
	MaxScore int16          `gorethink:"maxScore,omitempty" json:"maxScore"`
	Type     AssignmentType `gorethink:"type,omitempty" json:"type"`
	DueDate  time.Time      `gorethink:"dueDate,omitempty" json:"dueDate"`
	TimeStamp
}

// FieldMap ...
func (a *Assignment) FieldMap() binding.FieldMap {
	return binding.FieldMap{
		&a.ID:       "id",
		&a.Name:     "name",
		&a.TypeID:   "typeId",
		&a.ClassID:  "classId",
		&a.MaxScore: "maxScore",
		&a.TermID:   "termId",
		&a.DueDate:  "dueDate",
	}
}
func (a Assignment) Validate(req *http.Request, errs binding.Errors) binding.Errors {
	if a.Name == "" {
		errs = append(errs, RequiredErr("name"))
	}
	if a.ClassID == "" {
		errs = append(errs, RequiredErr("classId"))
	}
	if a.TermID == "" {
		errs = append(errs, RequiredErr("termId"))
	}
	if a.TypeID == "" {
		errs = append(errs, RequiredErr("typeId"))
	}
	if a.MaxScore <= 0 {
		errs = append(errs, RequiredErr("maxScore"))
	}
	return errs
}
