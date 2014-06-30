package model

import (
	"strings"
)

const (
	RequiredError = "RequiredError"
	// ContentTypeError     = "ContentTypeError"
	// DeserializationError = "DeserializationError"
	// TypeError            = "TypeError"
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

func (v *ValErrors) RequiredString(field, name string) {
	field = strings.TrimSpace(field)
	if field == "" {
		v.Add([]string{name}, RequiredError)
		return
	}
}
