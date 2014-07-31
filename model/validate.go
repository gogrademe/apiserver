package model

import (
	// "strings"

	"github.com/mholt/binding"
)

type (
	ValErrors []ValError
	ValError  struct {
		FieldNames []string `json:"fieldNames,omitempty"`
		Message    string   `json:"message,omitempty"`
	}
)

func (e ValError) Error() string {
	return e.Message
}

func (v *ValErrors) Len() int {
	return len(*v)
}

func (v *ValErrors) Add(fieldNames []string, message string) {
	*v = append(*v, ValError{
		FieldNames: fieldNames,
		Message:    message,
	})
}

func RequiredErr(fieldName string) binding.Error {
	return binding.Error{
		FieldNames: []string{fieldName},
		Message:    "missing required field",
	}
}
