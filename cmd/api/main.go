package main

import (
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/acekavi/keytide/internal/handlers"
	"github.com/acekavi/keytide/internal/server"
	"github.com/acekavi/keytide/pkg/middleware"
)

func main() {
	log := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	s := server.NewServer()

	// Method-and-wildcard patterns, which net/http's ServeMux has supported
	// since Go 1.22. A bare "/products" pattern matches every method and
	// everything beneath the path, so a DELETE to /products/1 reached a
	// handler written to list products. Fixing that was a common reason to
	// reach for a third-party router; there is no longer one here.
	s.Router.HandleFunc("GET /products", handlers.GetProducts)

	// The middleware was defined and never applied: the previous version
	// passed s.Router straight to ListenAndServe, so every request skipped
	// both. Order reads inside-out — logging is applied last, so it wraps auth
	// and therefore sees the 401s auth produces. Wrapping the other way round
	// would log only the requests that were already authorised.
	var handler http.Handler = s.Router
	handler = middleware.AuthMiddleware(middleware.RejectAll)(handler)
	handler = middleware.LoggingMiddleware(log)(handler)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: handler,
		// An http.Server with no timeouts holds a connection open indefinitely,
		// which is a denial of service that needs no attacker.
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       15 * time.Second,
		WriteTimeout:      15 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	log.Info("server starting", slog.String("addr", srv.Addr))
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Error("server stopped", slog.Any("error", err))
		os.Exit(1)
	}
}
