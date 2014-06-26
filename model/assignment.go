package model

import (
	"github.com/mholt/binding"

	"time"
)

type Assignment struct {
	ID        string    `gorethink:"id,omitempty"json:"id"`
	Name      string    `gorethink:"name"json:"name"`
	Type      string    `gorethink:"type"json:"tpe"`
	DueDate   time.Time `gorethink:"dueDate"json:"dueDate"`
	ClassTerm string    `gorethink:"classTerm"json:"classTerm"`
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
