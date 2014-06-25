package model

import (
	"time"
)

type Assignment struct {
	ID        string    `gorethink:"id,omitempty"json:"id"`
	Name      string    `gorethink:"name"json:"name"`
	Type      string    `gorethink:"type"json:"tpe"`
	DueDate   time.Time `gorethink:"dueDate"json:"dueDate"`
	ClassID   string    `gorethink:"classID"json:"classId"`
	ClassTerm string    `gorethink:"class_term_id"json:"classTerm"`
	TimeStamp
}
