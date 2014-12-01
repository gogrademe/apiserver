package model

import (
	"net/http"
	"time"

	"github.com/mholt/binding"
)

//Announcement ...
type Announcement struct {
	ID         string    `gorethink:"id,omitempty"json:"id"`
	Name       string    `gorethink:"name,omitempty"json:"name"`
	PersonID   string    `gorethink:"personId,omitempty"json:"personId"`
	PostedDate time.Time `gorethink:"postedDate,omitempty"json:"postedDate"`
	TimeStamp
}

// FieldMap ...
func (a *Announcement) FieldMap() binding.FieldMap {
	return binding.FieldMap{
		&a.ID:   "id",
		&a.Name: "name",
	}
}
func (a Announcement) Validate(req *http.Request, errs binding.Errors) binding.Errors {
	if a.Name == "" {
		errs = append(errs, RequiredErr("name"))
	}
	if a.PersonID == "" {
		errs = append(errs, RequiredErr("personId"))
	}
	return errs
}
