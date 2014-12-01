package model

import "time"

type Model interface {
	UpdateTime()
}

type TimeStamp struct {
	CreatedAt time.Time `gorethink:"createdAt"json:"createdAt"`
	UpdatedAt time.Time `gorethink:"updatedAt"json:"updatedAt"`
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
