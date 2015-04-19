package model

import "time"

type EmailConfirmation struct {
	ID     string    `gorethink:"id,omitempty" json:"id"`
	UserID string    `gorethink:"userId" json:"userId"`
	UsedOn time.Time `gorethink:"usedOn,omitempty" json:"usedOn"`
	TimeStamp
}
