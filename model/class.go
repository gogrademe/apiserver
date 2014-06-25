package model

import (
	"time"
)

type Class struct {
	ID         int64
	Name       string
	TeacherId  int64  `db:"teacher_id"`
	GradeLevel string `db:"grade_level"`
	Subject    string
	TimeStamp
}

type ClassTerm struct {
	ClassID   int64     `db:"class_id"`
	StartDate time.Time `db:"start_date"`
	EndDate   time.Time `db:"end_date"`
	TimeStamp
}
