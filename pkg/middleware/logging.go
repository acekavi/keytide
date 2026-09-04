package middleware

import (
	"log/slog"
	"net/http"
	"time"
)

// statusRecorder captures the status code on its way out.
//
// http.ResponseWriter does not expose what was written, so a logger that wants
// to report the status has to intercept WriteHeader. Embedding the interface
// leaves everything else — Header, Write — untouched.
type statusRecorder struct {
	http.ResponseWriter
	status int
}

func (r *statusRecorder) WriteHeader(code int) {
	r.status = code
	r.ResponseWriter.WriteHeader(code)
}

// LoggingMiddleware logs one structured line per request.
//
// It takes the logger rather than reaching for a package-level default, so a
// test can hand it a logger writing to a buffer and a deployment can choose
// JSON without this file knowing about it. The previous version used
// log.Printf to stderr — unparseable anywhere that collects logs — and
// reported only the method and path, not the status and not the duration.
func LoggingMiddleware(log *slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			// Default 200: a handler that writes a body without calling
			// WriteHeader has still sent 200, and the recorder must agree.
			rec := &statusRecorder{ResponseWriter: w, status: http.StatusOK}

			next.ServeHTTP(rec, r)

			log.InfoContext(r.Context(), "request",
				slog.String("method", r.Method),
				slog.String("path", r.URL.Path),
				slog.Int("status", rec.status),
				slog.Duration("took", time.Since(start)),
			)
		})
	}
}
