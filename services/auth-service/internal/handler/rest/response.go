package rest

import (
	"encoding/json"
	"net/http"
)

func sendResponse(w http.ResponseWriter, data any, statusCode int) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(data)
}
