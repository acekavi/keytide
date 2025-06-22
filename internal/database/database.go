package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// DBConfig holds database configuration
type DBConfig struct {
    Host     string
    Port     string
    User     string
    Password string
    DBName   string
    SSLMode  string
}

// NewPostgresDB creates and returns a new PostgreSQL database connection
func NewPostgresDB(config DBConfig) (*sql.DB, error) {
    // Build connection string
    connStr := fmt.Sprintf(
        "host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
        config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode,
    )

    // Open database connection
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        return nil, fmt.Errorf("failed to open database: %w", err)
    }

    // Test connection
    if err := db.Ping(); err != nil {
        return nil, fmt.Errorf("failed to ping database: %w", err)
    }

    // Initialize schema
    if err := initPostgresSchema(db); err != nil {
        return nil, fmt.Errorf("failed to initialize schema: %w", err)
    }

    // Seed with sample data if empty
    if err := SeedData(db); err != nil {
        log.Printf("Warning: Failed to seed database: %v", err)
    }

    return db, nil
}

// initSchema creates necessary tables if they don't exist
func initPostgresSchema(db *sql.DB) error {
    // Create products table
    _, err := db.Exec(`
        CREATE TABLE IF NOT EXISTS products (
            id TEXT PRIMARY KEY,
            name TEXT NOT NULL,
            price NUMERIC(10,2) NOT NULL,
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
        )
    `)
    if err != nil {
        return err
    }

    log.Println("PostgreSQL schema initialized successfully")
    return nil
}
