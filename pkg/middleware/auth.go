package middleware

import (
	"net/http"
	"strings"
)

// AuthMiddleware is a middleware for authenticating requests
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check for a valid token (this is just a simple example)
		token := r.Header.Get("Authorization")
		if !isValidToken(token) {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// isValidToken checks if the provided token is valid
func isValidToken(token string) bool {
	// In a real application, you would validate the token properly
	return strings.HasPrefix(token, "Bearer ")
}
