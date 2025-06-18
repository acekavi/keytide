package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/acekavi/keytide/internal/handlers"
	"github.com/acekavi/keytide/internal/server"
)

func main() {
    // Initialize the server
    s := server.NewServer()
    
    // Register routes
    s.Router.HandleFunc("/products", handlers.GetProducts)
    
    // Start server
    port := "8080"
    fmt.Printf("Server starting on port %s...\n", port)
    if err := http.ListenAndServe(":"+port, s.Router); err != nil {
        log.Fatal(err)
    }
}