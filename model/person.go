package model

import (
	"github.com/mholt/binding"
)

type (
	Person struct {
		ID         string    `gorethink:"id,omitempty"json:"id"`
		FirstName  string    `gorethink:"firstName,omitempty"json:"firstName,omitempty"`
		MiddleName string    `gorethink:"middleName,omitempty"json:"middleName,omitempty"`
		LastName   string    `gorethink:"lastName,omitempty"json:"lastName,omitempty"`
		Profiles   *Profiles `gorethink:"profiles,omitempty"json:"profiles,omitempty"`
		TimeStamp
	}
	Profiles struct {
		StudentID string `gorethink:"studentId"json:"studentId,omitempty"`
		TeacherID string `gorethink:"teacherId"json:"teacherId,omitempty"`
		ParentID  string `gorethink:"parentId"json:"parentId,omitempty"`
	}
)

func (p *Person) Validate() *ValErrors {
	var v *ValErrors

	v.RequiredString(p.FirstName, "firstName")
	v.RequiredString(p.LastName, "lastName")

	return v
}

func (p *Person) FieldMap() binding.FieldMap {
	return binding.FieldMap{
		&p.ID:         field("id", false),
		&p.FirstName:  field("firstName", true),
		&p.MiddleName: field("middleName", false),
		&p.LastName:   field("lastName", true),
	}
}
