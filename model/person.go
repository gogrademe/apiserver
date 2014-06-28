package model

import (
	"github.com/mholt/binding"
)

type Person struct {
	ID         string   `gorethink:"id,omitempty"json:"id"`
	FirstName  string   `gorethink:"firstName,omitempty"json:"firstName"`
	MiddleName string   `gorethink:"middleName,omitempty"json:"middleName"`
	LastName   string   `gorethink:"lastName,omitempty"json:"lastName"`
	Profiles   Profiles `gorethink:"profiles,omitempty"json:"profiles"`
	TimeStamp
}

type Profiles struct {
	StudentID string `gorethink:"studentId"json:"studentId"`
	TeacherID string `gorethink:"teacherId"json:"teacherId"`
	ParentID  string `gorethink:"parentId"json:"parentId"`
}

func (p *Person) FieldMap() binding.FieldMap {
	return binding.FieldMap{
		&p.ID:         field("id", false),
		&p.FirstName:  field("firstName", true),
		&p.MiddleName: field("middleName", false),
		&p.LastName:   field("lastName", true),
	}
}
