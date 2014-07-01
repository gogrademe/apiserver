package handlers

import (
	"encoding/json"
	"log"
	"net/http"
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

	// Context ...
	Context struct {
		Req    *http.Request
		Writer http.ResponseWriter
		Keys   map[string]interface{}
		index  int8
	}
)

// NotFoundErr should be used if a resource could not be found.
var NotFoundErr = &APIError{
	Code: 404,
}

const (
	serverError   = "server error"
	notFoundError = "not found"
)

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

// writeAPIError will write a JSON error to the client.
func writeAPIError(w http.ResponseWriter, apiErr *APIError) {
	if apiErr.Code == 0 {
		log.Println("ERROR CODE NOT SET USING 500", apiErr.Message, " Raw Error: ", apiErr.Raw)
		apiErr.Code = 500
	}
	log.Println("Friendly Message: ", apiErr.Message, " Raw Error: ", apiErr.Raw)
	e := &APIRes{
		"error": []APIError{
			*apiErr,
		},
	}
	writeJSONWithHeader(w, e, apiErr.Code)
	return
}

func writeJSON(w http.ResponseWriter, v *APIRes) {
	writeJSONWithHeader(w, v, 200)
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
