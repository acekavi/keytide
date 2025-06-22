package server

import (
	"github.com/acekavi/keytide/pkg/middleware"
	"github.com/gorilla/mux"
)

// Server represents the HTTP server
type Server struct {
    Router *mux.Router
}

// NewServer creates a new server instance
func NewServer() *Server {
    r := mux.NewRouter()
    
    // Apply global middlewares
    r.Use(middleware.CORS)
    
    return &Server{
        Router: r,
    }
}

// Group returns a new route group
func (s *Server) Group(path string, middlewares ...middleware.Middleware) *mux.Router {
    // Create a subrouter
    subRouter := s.Router.PathPrefix(path).Subrouter()
    
    // Apply group-specific middlewares
    for _, m := range middlewares {
        subRouter.Use(mux.MiddlewareFunc(m))
    }
    
    return subRouter
}