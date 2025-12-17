package helpers

import (
	"encoding/json"
	"net/http"
)

// APIResponse defines the standard JSON response
type APIResponse struct {
	Status  string      `json:"status"`            // "success" | "error"
	Message string      `json:"message,omitempty"` // optional message
	Data    interface{} `json:"data,omitempty"`    // payload
	Sound   string      `json:"sound,omitempty"`   // optional sound hint
}

// ValidationError represents a single validation error
type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// JSONSuccess writes a success JSON response
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

// JSONError writes an error JSON response
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

// JSONValidationError writes validation errors in a consistent format
func JSONValidationError(w http.ResponseWriter, errors []ValidationError, code int) {
	res := APIResponse{
		Status: "error",
		Data:   errors,
		Sound:  "error_sound",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(res)
}

// DecodeJSON decodes request body into the given struct, preventing unknown fields
func DecodeJSON(r *http.Request, dest interface{}) error {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	return decoder.Decode(dest)
}
