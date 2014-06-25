package model

import (
	"time"
)

type Class struct {
	ID         string `gorethink:"id,omitempty"`
	Name       string	`gorethink:"name"`
	// TeacherId  string  `gorethink:"teacher_id"`
	GradeLevel string `gorethink:"grade_level"`
	Subject    string
	TimeStamp
}

type ClassTerm struct {
	ClassID   int64     `gorethink:"class_id"`
	StartDate time.Time `gorethink:"start_date"`
	EndDate   time.Time `gorethink:"end_date"`
	TimeStamp
}
