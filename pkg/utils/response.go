package utils

import (
	"encoding/json"
	"net/http"
)

// JSONResponse sends a JSON response with the given status code and data
func JSONResponse(w http.ResponseWriter, statusCode int, data any) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(statusCode)
    json.NewEncoder(w).Encode(data)
}

// ErrorResponse represents an error response
type ErrorResponse struct {
    Error string `json:"error"`
}

// JSONError sends a JSON error response with the given status code and error message
func JSONError(w http.ResponseWriter, statusCode int, message string) {
    JSONResponse(w, statusCode, ErrorResponse{Error: message})
}