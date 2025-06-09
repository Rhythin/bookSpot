package v1

import (
	"encoding/json"
	"net/http"
)

// sendResponse uses the response writer to send a response with the given data and status code
func sendResponse(w http.ResponseWriter, data any, code int) (err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if data != nil {
		return json.NewEncoder(w).Encode(data)
	}
	return nil
}
