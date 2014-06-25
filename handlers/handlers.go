package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// APIError represents an error produced by the API
type APIError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type APIRes map[string]interface{}

const serverError = "server error"

// writeError will write a JSON error to the client.
func writeError(w http.ResponseWriter, message string, code int) {
	e := APIError{
		Code:    code,
		Message: message,
	}
	data, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(code)
	w.Write(data)
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
