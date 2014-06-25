package model

import (
	"time"
)

type TimeStamp struct {
	CreatedAt time.Time `db:"created_at"json:"createdAt"`
	UpdatedAt time.Time `db:"updated_at"json:"updatedAt"`
}

func (a *TimeStamp) UpdateTime() {

	t := time.Now().UTC()
	if !a.CreatedAt.IsZero() {
		a.UpdatedAt = t
		return
	}
	a.CreatedAt = t
	a.UpdatedAt = t
	return
}
