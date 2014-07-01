package model

import (
	// "net/http"

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

// func (s Student) Validate(req *http.Request, errs binding.Errors) binding.Errors {
// 	if s.Message == "Go needs generics" {
// 		errs = append(errs, binding.Error{
// 			FieldNames:     []string{"message"},
// 			Classification: "ComplaintError",
// 			Message:        "Go has generics. They're called interfaces.",
// 		})
// 	}
// 	return errs
// }

func (s *Student) FieldMap() binding.FieldMap {
	return binding.FieldMap{
		&s.ID:         field("id", false),
		&s.PersonID:   field("personId", true),
		&s.GradeLevel: field("gradeLevel", true),
	}
}
