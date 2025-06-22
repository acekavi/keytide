package middleware

import (
	"net/http"
)

// Middleware defines a function to process HTTP requests
type Middleware func(http.Handler) http.Handler

// Chain applies middlewares to a http.Handler
func Chain(h http.Handler, middlewares ...Middleware) http.Handler {
    for _, m := range middlewares {
        h = m(h)
    }
    return h
}