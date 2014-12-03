package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gogrademe/apiserver/store"
)

type (

	// APIRes response from the API.
	APIRes map[string]interface{}

	// APIError represents an error produced by the API
	APIError struct {
		Code    int         `json:"code"`
		Type    string      `json:"type"`
		Message interface{} `json:"message"`
		Raw     string      `json:"-"`
	}
)

// NotFoundErr should be used if a resource could not be found.
var NotFoundErr = &APIError{
	Code: 404,
}

const (
	StatusUnprocessable = 422
	serverError         = "server error"
	notFoundError       = "not found"
)

func handleDBError(rw http.ResponseWriter, err error) {
	if err == store.ErrNotFound {
		rw.WriteHeader(http.StatusNotFound)
		// writeError(c.Writer, notFoundError, 404, nil)
		return
	}
	if err != nil {
		writeError(c.Writer, serverError, 500, nil)
		return
	}
}

// writeError will write a JSON error to the client.
func writeError(w http.ResponseWriter, message interface{}, code int, errorToLog error) {
	log.Println("Friendly Message: ", message, " Raw Error: ", errorToLog)
	e := &APIRes{
		"error": []APIError{
			APIError{
				Code:    code,
				Message: message,
			},
		},
	}
	writeJSONWithHeader(w, e, code)
	return
}

func writeJSONWithHeader(w http.ResponseWriter, v *APIRes, code int) {
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		log.Printf("Error marshalling json: %v", err)
		w.WriteHeader(500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
	return
}
