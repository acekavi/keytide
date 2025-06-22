package main

import (
	"database/sql"
	"log"
	"net/http"

	config "github.com/acekavi/keytide/configs"
	"github.com/acekavi/keytide/internal/database"
	"github.com/acekavi/keytide/internal/handlers"
	"github.com/acekavi/keytide/internal/repository"
	"github.com/gorilla/mux"
)

func main() {
    // Load configuration
    cfg := config.LoadConfig()

    var db *sql.DB
    var err error

    dbConfig := database.DBConfig{
        Host:     cfg.Database.Host,
        Port:     cfg.Database.Port,
        User:     cfg.Database.User,
        Password: cfg.Database.Password,
        DBName:   cfg.Database.DBName,
        SSLMode:  cfg.Database.SSLMode,
    }

    log.Printf("Using PostgreSQL database at: %s:%s/%s", dbConfig.Host, dbConfig.Port, dbConfig.DBName)
    db, err = database.NewPostgresDB(dbConfig)

    if err != nil {
        log.Fatalf("Failed to initialize database: %v", err)
    }
    defer db.Close()

    // Initialize repository
    productRepo := repository.NewDBProductRepository(db)

    // Initialize handlers
    productHandler := handlers.NewProductHandler(productRepo)

    // Initialize router
    r := mux.NewRouter()

    // Register routes
    r.HandleFunc("/products", productHandler.GetProducts).Methods("GET")
    r.HandleFunc("/products/{id}", productHandler.GetProduct).Methods("GET")
    r.HandleFunc("/products", productHandler.CreateProduct).Methods("POST")
    r.HandleFunc("/products/{id}", productHandler.UpdateProduct).Methods("PUT")
    r.HandleFunc("/products/{id}", productHandler.DeleteProduct).Methods("DELETE")

    // Start server
    log.Printf("Server starting on :%s", cfg.Server.Port)
    if err := http.ListenAndServe(":"+cfg.Server.Port, r); err != nil {
        log.Fatal(err)
    }
}
