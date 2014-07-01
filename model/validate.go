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

// func (v *ValErrors) RequiredString(field, name string) {
// 	field = strings.TrimSpace(field)
// 	if field == "" {
// 		v.Add([]string{name}, RequiredError)
// 		return
// 	}
// }
func RequiredErr(fieldName string) binding.Error {
	return binding.Error{
		FieldNames: []string{fieldName},
		Message:    "missing required field",
	}
}

// func RequiredString(errs binding.Errors, fieldName string, value string) {
// 	value = strings.TrimSpace(value)
// 	if value == "" {
// 		errs = append(errs, binding.Error{
// 			FieldNames: []string{fieldName},
// 			Message:    "missing required field",
// 		})
// 	}
// }
