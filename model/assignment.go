package model

import (
	"time"

	"github.com/mholt/binding"
)

type Assignment struct {
	ID        string    `gorethink:"id,omitempty"json:"id"`
	Name      string    `gorethink:"name,omitempty"json:"name"`
	Type      string    `gorethink:"type,omitempty"json:"type"`
	DueDate   time.Time `gorethink:"dueDate,omitempty"json:"dueDate"`
	ClassTerm string    `gorethink:"classTerm,omitempty"json:"classTerm"`
	TimeStamp
}

func (a *Assignment) FieldMap() binding.FieldMap {
	return binding.FieldMap{
		&a.ID:        field("id", false),
		&a.Name:      field("name", true),
		&a.Type:      field("type", true),
		&a.DueDate:   field("dueDate", true),
		&a.ClassTerm: field("classTerm", true),
	}
}

func (a *Assignment) Validate() *ValErrors {
	var v *ValErrors

	v.RequiredString(a.Name, "name")
	v.RequiredString(a.Type, "type")
	// v.RequiredString(a.DueDate, "id")
	v.RequiredString(a.ClassTerm, "classTerm")

	return v
}
