package repository

import (
	"fmt"

	"keytide/internal/models"
)

// ProductRepository defines methods for product data access
type ProductRepository interface {
    GetAll() ([]models.Product, error)
    GetByID(id string) (models.Product, error)
    Create(product models.Product) error
    Update(product models.Product) error
    Delete(id string) error
}

// InMemoryProductRepository implements ProductRepository with an in-memory data store
type InMemoryProductRepository struct {
    products map[string]models.Product
}

// NewInMemoryProductRepository creates a new in-memory product repository
func NewInMemoryProductRepository() *InMemoryProductRepository {
    return &InMemoryProductRepository{
        products: map[string]models.Product{
            "1": {ID: "1", Name: "Laptop", Price: 999.99},
            "2": {ID: "2", Name: "Headphones", Price: 99.99},
        },
    }
}

// GetAll returns all products
func (r *InMemoryProductRepository) GetAll() ([]models.Product, error) {
    products := make([]models.Product, 0, len(r.products))
    for _, product := range r.products {
        products = append(products, product)
    }
    return products, nil
}

// GetByID returns a product by ID
func (r *InMemoryProductRepository) GetByID(id string) (models.Product, error) {
    product, exists := r.products[id]
    if !exists {
        return models.Product{}, fmt.Errorf("product with ID %s not found", id)
    }
    return product, nil
}

// Create adds a new product
func (r *InMemoryProductRepository) Create(product models.Product) error {
    if _, exists := r.products[product.ID]; exists {
        return fmt.Errorf("product with ID %s already exists", product.ID)
    }
    r.products[product.ID] = product
    return nil
}

// Update modifies an existing product
func (r *InMemoryProductRepository) Update(product models.Product) error {
    if _, exists := r.products[product.ID]; !exists {
        return fmt.Errorf("product with ID %s not found", product.ID)
    }
    r.products[product.ID] = product
    return nil
}

// Delete removes a product
func (r *InMemoryProductRepository) Delete(id string) error {
    if _, exists := r.products[id]; !exists {
        return fmt.Errorf("product with ID %s not found", id)
    }
    delete(r.products, id)
    return nil
}