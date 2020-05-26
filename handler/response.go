package handler

import (
	"encoding/json"
	"net/http"
)

// Serve data as JSON as response
func jsonResponse(w http.ResponseWriter, status int, v interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}
