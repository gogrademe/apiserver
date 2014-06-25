package model

import (
	"time"
)

type Assignment struct {
	ID        string    `gorethink:"id,omitempty"`
	Name      string    `gorethink:"name"`
	Type      string    `gorethink:"type"`
	DueDate   time.Time `gorethink:"dueDate"`
	ClassID   string    `gorethink:"classID"`
	ClassTerm string    `gorethink:"class_term_id"`
	TimeStamp
}
