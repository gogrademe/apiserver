package model

import (
	"github.com/mholt/binding"
)

type Student struct {
	ID         string `gorethink:"id,omitempty"json:"id"`
	PersonID   string `gorethink:"personId,omitempty"json:"personId"`
	GradeLevel string `gorethink:"gradeLevel,omitempty"json:"gradeLevel"`
	TimeStamp
}

func (s *Student) Validate() *ValErrors {
	var v *ValErrors

	v.RequiredString(s.PersonID, "personId")
	v.RequiredString(s.GradeLevel, "gradeLevel")

	return v
}

func (s *Student) FieldMap() binding.FieldMap {
	return binding.FieldMap{
		&s.ID:         field("id", false),
		&s.PersonID:   field("personId", true),
		&s.GradeLevel: field("gradeLevel", true),
	}
}
