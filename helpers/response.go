package helpers

import (
	"encoding/json"
	"net/http"
)

type APIResponse struct {
	Status  string      `json:"status"`            // success | error
	Message string      `json:"message,omitempty"` // optional error or success message
	Data    interface{} `json:"data,omitempty"`    // payload
	Sound   string      `json:"sound,omitempty"`   // optional sound hint
}

// Success writes a success JSON response
func JSONSuccess(w http.ResponseWriter, data interface{}, message string, sound string) {
	res := APIResponse{
		Status:  "success",
		Message: message,
		Data:    data,
		Sound:   sound,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

// Error writes an error JSON response
func JSONError(w http.ResponseWriter, message string, sound string, code int) {
	res := APIResponse{
		Status:  "error",
		Message: message,
		Sound:   sound,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(res)
}
