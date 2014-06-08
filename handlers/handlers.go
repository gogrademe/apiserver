package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

// Error represents an error produced by the API
type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// ErrorResponse represents the JSON response for endpoints which only return an error
type ErrorResponse struct {
	Error  *Error              `json:"error"`
	writer http.ResponseWriter `json:"-"`
}

func (e *ErrorResponse) ReturnError(code int, message string) {
	writeJson(e.writer, ErrorResponse{
		Error: &Error{
			Code:    code,
			Message: message,
		},
	})
}

/* https://github.com/DenverGophers/talks/blob/master/2013-04/mgo/example_6/read_json.go */

// readJson will parses the JSON-encoded data in the http request and store the result in v
func readJson(r *http.Request, v interface{}) bool {
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
func writeJson(w http.ResponseWriter, v interface{}) {
	// if data, err := json.Marshal(v); err != nil {
	if data, err := json.MarshalIndent(v, "", "  "); err != nil {
		log.Printf("Error marshalling json: %v", err)
	} else {
		w.Header().Set("Content-Length", strconv.Itoa(len(data)))
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	}
}
