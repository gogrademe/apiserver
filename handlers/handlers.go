package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// APIError represents an error produced by the API
type APIError struct {
	Code    int         `json:"code"`
	Message interface{} `json:"message"`
}

type APIRes map[string]interface{}

const serverError = "server error"
const notFoundError = "not found"

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
	data, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
	return
}
func writeJSON(w http.ResponseWriter, v *APIRes) {
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		log.Printf("Error marshalling json: %v", err)
		w.WriteHeader(500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
	return
}

/* https://github.com/DenverGophers/talks/blob/master/2013-04/mgo/example_6/read_json.go */

// readJson will parses the JSON-encoded data in the http request and store the result in v
func readJSON(r *http.Request, v interface{}) bool {
	defer r.Body.Close()

	var (
		body []byte
		err  error
	)

	body, err = ioutil.ReadAll(r.Body)

	if err != nil {
		log.Printf("ReadJson couldn't read request body %v", err)
		return false
	}

	if err = json.Unmarshal(body, v); err != nil {
		log.Printf("ReadJson couldn't parse request body %v", err)
		return false
	}

	return true
}
