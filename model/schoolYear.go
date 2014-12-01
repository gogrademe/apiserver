package model

import (
	"net/http"

	"github.com/mholt/binding"
)

type SchoolYear struct {
	ID    string `gorethink:"id,omitempty"json:"id"`
	Start int    `gorethink:"start,omitempty"json:"start"`
	End   int    `gorethink:"end,omitempty"json:"end"`
	TimeStamp
}

func (t SchoolYear) Validate(req *http.Request, errs binding.Errors) binding.Errors {
	if t.Start == 0 {
		errs = append(errs, RequiredErr("start"))
	}
	if t.End == 0 {
		errs = append(errs, RequiredErr("end"))
	}
	return errs
}

func (t *SchoolYear) FieldMap() binding.FieldMap {
	return binding.FieldMap{
		&t.ID:    "id",
		&t.Start: "start",
		&t.End:   "end",
	}
}
