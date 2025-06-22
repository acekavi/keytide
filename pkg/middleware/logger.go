package middleware

import (
	"net/http"
	"time"

	"github.com/acekavi/keytide/pkg/logger"
	"go.uber.org/zap"
)

// Logger logs HTTP requests
func Logger(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        
        // Create a custom response writer to capture status code
        rw := &responseWriter{w, http.StatusOK}
        
        // Call the next handler
        next.ServeHTTP(rw, r)
        
        // Log the request with structured logging
        duration := time.Since(start)
        logger.Info("HTTP Request",
            zap.String("method", r.Method),
            zap.String("path", r.URL.Path),
            zap.Int("status", rw.statusCode),
            zap.Duration("duration", duration),
            zap.String("remote", r.RemoteAddr),
            zap.String("user-agent", r.UserAgent()),
        )
    })
}

// responseWriter is a wrapper for http.ResponseWriter that captures the status code
type responseWriter struct {
    http.ResponseWriter
    statusCode int
}

// WriteHeader captures the status code before writing it
func (rw *responseWriter) WriteHeader(code int) {
    rw.statusCode = code
    rw.ResponseWriter.WriteHeader(code)
}