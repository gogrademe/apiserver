package model

import (
	"github.com/mholt/binding"
)

type Student struct {
	ID         string `gorethink:"id,omitempty"json:"id"`
	PersonID   string `gorethink:"personID"json:"personId"`
	GradeLevel string `gorethink:"gradeLevel"json:"gradeLevel"`
	TimeStamp
}

func (s *Student) FieldMap() binding.FieldMap {
	return binding.FieldMap{
		&s.ID:         field("id", false),
		&s.PersonID:   field("personId", true),
		&s.GradeLevel: field("gradeLevel", true),
	}
}
