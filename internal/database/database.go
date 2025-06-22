package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

// NewSQLiteDB creates and returns a new SQLite database connection
func NewSQLiteDB(dbPath string) (*sql.DB, error) {
    // Ensure directory exists
    dir := filepath.Dir(dbPath)
    if err := os.MkdirAll(dir, 0755); err != nil {
        return nil, fmt.Errorf("failed to create database directory: %w", err)
    }

    // Open database connection
    db, err := sql.Open("sqlite3", dbPath)
    if err != nil {
        return nil, fmt.Errorf("failed to open database: %w", err)
    }

    // Test connection
    if err := db.Ping(); err != nil {
        return nil, fmt.Errorf("failed to ping database: %w", err)
    }

    // Initialize schema
    if err := initSchema(db); err != nil {
        return nil, fmt.Errorf("failed to initialize schema: %w", err)
    }

    return db, nil
}

// initSchema creates necessary tables if they don't exist
func initSchema(db *sql.DB) error {
    // Create products table
    _, err := db.Exec(`
        CREATE TABLE IF NOT EXISTS products (
            id TEXT PRIMARY KEY,
            name TEXT NOT NULL,
            price REAL NOT NULL,
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
        )
    `)
    if err != nil {
        return err
    }

    log.Println("Database schema initialized successfully")
    return nil
}