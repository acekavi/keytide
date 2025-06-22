package middleware

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"reflect"

	"github.com/acekavi/keytide/pkg/utils"
	"github.com/acekavi/keytide/pkg/validator"
)

// ValidateRequest validates request body against a model
func ValidateRequest(model interface{}) func(next http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            // Create a new instance of the model
            val := reflect.New(reflect.TypeOf(model)).Interface()
            
            // Read body
            body, err := io.ReadAll(r.Body)
            if err != nil {
                utils.JSONError(w, http.StatusBadRequest, "Failed to read request body")
                return
            }
            
            // Restore body for later use
            r.Body = io.NopCloser(bytes.NewBuffer(body))
            
            // Decode JSON
            if err := json.Unmarshal(body, val); err != nil {
                utils.JSONError(w, http.StatusBadRequest, "Invalid JSON format")
                return
            }
            
            // Validate
            if errors := validator.Validate(val); len(errors) > 0 {
                utils.JSONResponse(w, http.StatusBadRequest, map[string]any{
                    "errors": errors,
                })
                return
            }
            
            // Call next handler
            next.ServeHTTP(w, r)
        })
    }
}