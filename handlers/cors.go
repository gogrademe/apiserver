package handlers

import (
	"net/http"
)

func CORS(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	// w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token")
	// w.Header().Set("Access-Control-Allow-Credentials", "true")

	return
}

// func AddDefaultHeaders(fn http.HandlerFunc) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {

// 		w.Header().Set("Access-Control-Allow-Origin", "*")

// 		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
// 		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token")
// 		w.Header().Set("Access-Control-Allow-Credentials", "true")
// 		fn(w, r)
// 	}
// }
