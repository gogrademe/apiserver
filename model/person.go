package model

import (
	"github.com/mholt/binding"
)

type Person struct {
	ID         string   `gorethink:"id,omitempty"json:"id"`
	FirstName  string   `gorethink:"firstName"json:"firstName"`
	MiddleName string   `gorethink:"middleName"json:"middleName"`
	LastName   string   `gorethink:"lastName"json:"lastName"`
	Profiles   Profiles `gorethink:"profiles"json:"profiles"`
	TimeStamp
}

type Profiles struct {
	StudentID string `gorethink:"studentId"json:"studentId"`
	TeacherID string `gorethink:"teacherId"json:"teacherId"`
	ParentID  string `gorethink:"parentId"json:"parentId"`
}

// map[string]string `gorethink:"profiles,omitempty"json:"profiles,omitempty"`

func (p *Person) FieldMap() binding.FieldMap {
	return binding.FieldMap{
		&p.ID:         field("id", false),
		&p.FirstName:  field("firstName", true),
		&p.MiddleName: field("middleName", false),
		&p.LastName:   field("lastName", true),
	}
}
