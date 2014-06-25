package model

import (
	"time"
)

type Class struct {
	ID         string `gorethink:"id,omitempty"json:"id"`
	Name       string `gorethink:"name"json:"name"`
	GradeLevel string `gorethink:"gradeLevel"json:"gradeLevel"`
	Subject    string `gorethink:"subject"json:"subject"`
	Terms      []string
	TimeStamp
}

type ClassTerm struct {
	// ClassID   int64     `gorethink:"classId"json:"classId"`
	StartDate time.Time `gorethink:"startDate"json:"startDate"`
	EndDate   time.Time `gorethink:"endDate"json:"endDate"`
	TimeStamp
}
