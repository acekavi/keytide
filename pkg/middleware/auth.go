package middleware

import (
	"net/http"

	"github.com/acekavi/keytide/pkg/utils"
)

// Authenticate checks for a valid API key
func Authenticate(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Get API key from header
        apiKey := r.Header.Get("X-API-Key")
        
        // For demonstration, use a hard-coded API key
        // In production, you would validate against a database or auth service
        if apiKey == "" || apiKey != "your-api-key" {
            utils.JSONError(w, http.StatusUnauthorized, "Invalid or missing API key")
            return
        }
        
        next.ServeHTTP(w, r)
    })
}