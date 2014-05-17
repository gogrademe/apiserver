package models

import (
	"time"
)

type Teacher struct {
	Id         int64
	FirstName  string
	MiddleName string
	LastName   string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
