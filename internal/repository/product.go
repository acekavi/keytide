package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/acekavi/keytide/internal/models"
)

// ProductRepository defines methods for product data access
type ProductRepository interface {
    GetAll() ([]models.Product, error)
    GetByID(id string) (models.Product, error)
    Create(product models.Product) error
    Update(product models.Product) error
    Delete(id string) error
}

// DBProductRepository implements ProductRepository with a database
type DBProductRepository struct {
    db *sql.DB
}

// NewDBProductRepository creates a new database product repository
func NewDBProductRepository(db *sql.DB) *DBProductRepository {
    return &DBProductRepository{
        db: db,
    }
}

// GetAll returns all products
func (r *DBProductRepository) GetAll() ([]models.Product, error) {
    rows, err := r.db.Query("SELECT id, name, price FROM products")
    if err != nil {
        return nil, fmt.Errorf("failed to query products: %w", err)
    }
    defer rows.Close()

    var products []models.Product
    for rows.Next() {
        var p models.Product
        if err := rows.Scan(&p.ID, &p.Name, &p.Price); err != nil {
            return nil, fmt.Errorf("failed to scan product: %w", err)
        }
        products = append(products, p)
    }

    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("error iterating products: %w", err)
    }

    return products, nil
}

// GetByID returns a product by ID
func (r *DBProductRepository) GetByID(id string) (models.Product, error) {
    var p models.Product
    err := r.db.QueryRow("SELECT id, name, price FROM products WHERE id = ?", id).
        Scan(&p.ID, &p.Name, &p.Price)
    if err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            return models.Product{}, fmt.Errorf("product with ID %s not found", id)
        }
        return models.Product{}, fmt.Errorf("failed to query product: %w", err)
    }
    return p, nil
}

// Create adds a new product
func (r *DBProductRepository) Create(product models.Product) error {
    _, err := r.db.Exec(
        "INSERT INTO products (id, name, price) VALUES (?, ?, ?)",
        product.ID, product.Name, product.Price,
    )
    if err != nil {
        return fmt.Errorf("failed to create product: %w", err)
    }
    return nil
}

// Update modifies an existing product
func (r *DBProductRepository) Update(product models.Product) error {
    result, err := r.db.Exec(
        "UPDATE products SET name = ?, price = ? WHERE id = ?",
        product.Name, product.Price, product.ID,
    )
    if err != nil {
        return fmt.Errorf("failed to update product: %w", err)
    }

    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return fmt.Errorf("failed to get rows affected: %w", err)
    }

    if rowsAffected == 0 {
        return fmt.Errorf("product with ID %s not found", product.ID)
    }

    return nil
}

// Delete removes a product
func (r *DBProductRepository) Delete(id string) error {
    result, err := r.db.Exec("DELETE FROM products WHERE id = ?", id)
    if err != nil {
        return fmt.Errorf("failed to delete product: %w", err)
    }

    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return fmt.Errorf("failed to get rows affected: %w", err)
    }

    if rowsAffected == 0 {
        return fmt.Errorf("product with ID %s not found", id)
    }

    return nil
}