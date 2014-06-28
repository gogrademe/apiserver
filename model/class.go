package model

import (
	"time"

	"github.com/mholt/binding"
)

type Class struct {
	ID         string            `gorethink:"id,omitempty"json:"id"`
	Name       string            `gorethink:"name"json:"name"`
	GradeLevel string            `gorethink:"gradeLevel"json:"gradeLevel"`
	Subject    string            `gorethink:"subject"json:"subject"`
	Terms      map[string]string `gorethink:"-"json:"terms"`
	TimeStamp
}

func (c *Class) FieldMap() binding.FieldMap {
	return binding.FieldMap{
		&c.ID:         field("id", false),
		&c.Name:       field("name", true),
		&c.GradeLevel: field("gradeLevel", false),
		&c.Subject:    field("subject", true),
	}
}

type ClassTerm struct {
	ID          string       `gorethink:"id,omitempty"json:"id"`
	ClassID     string       `gorethink:"classId,omitempty"json:"classId"`
	StartDate   time.Time    `gorethink:"startDate"json:"startDate"`
	EndDate     time.Time    `gorethink:"endDate"json:"endDate"`
	Assignments []Assignment `gorethink:"-"json:"assignments"`
	People      []Person     `gorethink:"-"json:"people"`
	TimeStamp
}

func (ct *ClassTerm) FieldMap() binding.FieldMap {
	return binding.FieldMap{
		&ct.ID:        field("id", false),
		&ct.ClassID:   field("classId", true),
		&ct.StartDate: field("startDate", false),
		&ct.EndDate:   field("endDate", false),
	}
}
