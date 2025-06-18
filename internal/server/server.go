package server

import (
	"net/http"
)

// Server represents the HTTP server
type Server struct {
    Router *http.ServeMux
}

// NewServer creates a new server instance
func NewServer() *Server {
    return &Server{
        Router: http.NewServeMux(),
    }
}