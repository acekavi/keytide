package database

import (
	"database/sql"
	"fmt"
	"log"
)

// SeedData adds initial data to the database if it's empty
func SeedData(db *sql.DB) error {
    // Check if products table is empty
    var count int
    err := db.QueryRow("SELECT COUNT(*) FROM products").Scan(&count)
    if err != nil {
        return fmt.Errorf("failed to check if products table is empty: %w", err)
    }

    // If not empty, don't seed
    if count > 0 {
        return nil
    }

    // Sample products
    products := []struct {
        ID    string
        Name  string
        Price float64
    }{
        {"1", "Laptop", 999.99},
        {"2", "Headphones", 99.99},
        {"3", "Keyboard", 129.99},
    }

    // Insert sample data
    for _, p := range products {
        _, err = db.Exec(
                "INSERT INTO products (id, name, price) VALUES ($1, $2, $3)",
                p.ID, p.Name, p.Price,
            )
        if err != nil {
            return fmt.Errorf("failed to seed product data: %w", err)
        }
    }

    log.Println("Database seeded with sample data")
    return nil
}
