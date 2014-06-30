package model

import (
	"time"

	"github.com/mholt/binding"
)

//Assignment ...
type Assignment struct {
	ID      string    `gorethink:"id,omitempty"json:"id"`
	TermID  string    `gorethink:"termId,omitempty"json:"termId"`
	Name    string    `gorethink:"name,omitempty"json:"name"`
	Type    string    `gorethink:"type,omitempty"json:"type"`
	DueDate time.Time `gorethink:"dueDate,omitempty"json:"dueDate"`
	TimeStamp
}

// FieldMap ...
func (a *Assignment) FieldMap() binding.FieldMap {
	return binding.FieldMap{
		&a.ID:      field("id", false),
		&a.TermID:  field("termId", true),
		&a.Name:    field("name", true),
		&a.Type:    field("type", true),
		&a.DueDate: field("dueDate", true),
	}
}

// Validate ...
func (a *Assignment) Validate() *ValErrors {
	var v *ValErrors

	v.RequiredString(a.Name, "name")
	v.RequiredString(a.Type, "type")
	// v.RequiredString(a.DueDate, "id")

	return v
}
