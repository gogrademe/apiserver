package model

import (
	"time"

	"github.com/mholt/binding"
)

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

func field(form string, required bool) binding.Field {
	return binding.Field{
		Form:     form,
		Required: required,
	}
}
