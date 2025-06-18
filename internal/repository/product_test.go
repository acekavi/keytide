package repository

import (
	"testing"

	"github.com/acekavi/keytide/internal/models"
)

func TestInMemoryProductRepository_GetAll(t *testing.T) {
    // Create a new repository with some initial data
    repo := NewInMemoryProductRepository()
    
    // Get all products
    products, err := repo.GetAll()
    
    // Check if there was an error
    if err != nil {
        t.Errorf("Expected no error, got %v", err)
    }
    
    // Check if we got the expected number of products
    if len(products) != 2 {
        t.Errorf("Expected 2 products, got %d", len(products))
    }
    
    // Verify the products exist in the result
    found1, found2 := false, false
    for _, p := range products {
        if p.ID == "1" && p.Name == "Laptop" && p.Price == 999.99 {
            found1 = true
        }
        if p.ID == "2" && p.Name == "Headphones" && p.Price == 99.99 {
            found2 = true
        }
    }
    
    if !found1 {
        t.Error("Expected to find product with ID 1")
    }
    if !found2 {
        t.Error("Expected to find product with ID 2")
    }
}

func TestInMemoryProductRepository_GetByID(t *testing.T) {
    // Create a new repository with some initial data
    repo := NewInMemoryProductRepository()
    
    // Test cases
    testCases := []struct {
        name        string
        id          string
        expectError bool
        expected    models.Product
    }{
        {
            name:        "existing product",
            id:          "1",
            expectError: false,
            expected:    models.Product{ID: "1", Name: "Laptop", Price: 999.99},
        },
        {
            name:        "non-existing product",
            id:          "999",
            expectError: true,
            expected:    models.Product{},
        },
    }
    
    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            product, err := repo.GetByID(tc.id)
            
            // Check error expectation
            if tc.expectError && err == nil {
                t.Error("Expected an error but didn't get one")
            }
            if !tc.expectError && err != nil {
                t.Errorf("Didn't expect an error but got: %v", err)
            }
            
            // If we don't expect an error, check the product
            if !tc.expectError {
                if product.ID != tc.expected.ID {
                    t.Errorf("Expected ID %s, got %s", tc.expected.ID, product.ID)
                }
                if product.Name != tc.expected.Name {
                    t.Errorf("Expected Name %s, got %s", tc.expected.Name, product.Name)
                }
                if product.Price != tc.expected.Price {
                    t.Errorf("Expected Price %.2f, got %.2f", tc.expected.Price, product.Price)
                }
            }
        })
    }
}

func TestInMemoryProductRepository_Create(t *testing.T) {
    // Create a new repository
    repo := NewInMemoryProductRepository()
    
    // Create a new product
    newProduct := models.Product{
        ID:    "3",
        Name:  "Mouse",
        Price: 29.99,
    }
    
    // Test create
    err := repo.Create(newProduct)
    if err != nil {
        t.Errorf("Failed to create product: %v", err)
    }
    
    // Verify product was created
    product, err := repo.GetByID("3")
    if err != nil {
        t.Errorf("Failed to get created product: %v", err)
    }
    
    if product.ID != newProduct.ID {
        t.Errorf("Expected ID %s, got %s", newProduct.ID, product.ID)
    }
    if product.Name != newProduct.Name {
        t.Errorf("Expected Name %s, got %s", newProduct.Name, product.Name)
    }
    if product.Price != newProduct.Price {
        t.Errorf("Expected Price %.2f, got %.2f", newProduct.Price, product.Price)
    }
    
    // Test creating product with existing ID
    err = repo.Create(models.Product{ID: "1", Name: "Duplicate", Price: 10.99})
    if err == nil {
        t.Error("Expected error when creating product with duplicate ID, but got none")
    }
}

func TestInMemoryProductRepository_Update(t *testing.T) {
    // Create a new repository
    repo := NewInMemoryProductRepository()
    
    // Update an existing product
    updatedProduct := models.Product{
        ID:    "1",
        Name:  "Updated Laptop",
        Price: 1299.99,
    }
    
    err := repo.Update(updatedProduct)
    if err != nil {
        t.Errorf("Failed to update product: %v", err)
    }
    
    // Verify product was updated
    product, err := repo.GetByID("1")
    if err != nil {
        t.Errorf("Failed to get updated product: %v", err)
    }
    
    if product.Name != updatedProduct.Name {
        t.Errorf("Expected Name %s, got %s", updatedProduct.Name, product.Name)
    }
    if product.Price != updatedProduct.Price {
        t.Errorf("Expected Price %.2f, got %.2f", updatedProduct.Price, product.Price)
    }
    
    // Test updating non-existent product
    err = repo.Update(models.Product{ID: "999", Name: "Non-existent", Price: 9.99})
    if err == nil {
        t.Error("Expected error when updating non-existent product, but got none")
    }
}

func TestInMemoryProductRepository_Delete(t *testing.T) {
    // Create a new repository
    repo := NewInMemoryProductRepository()
    
    // Delete an existing product
    err := repo.Delete("1")
    if err != nil {
        t.Errorf("Failed to delete product: %v", err)
    }
    
    // Verify product was deleted
    _, err = repo.GetByID("1")
    if err == nil {
        t.Error("Expected error when getting deleted product, but got none")
    }
    
    // Check number of products
    products, _ := repo.GetAll()
    if len(products) != 1 {
        t.Errorf("Expected 1 product after deletion, got %d", len(products))
    }
    
    // Test deleting non-existent product
    err = repo.Delete("999")
    if err == nil {
        t.Error("Expected error when deleting non-existent product, but got none")
    }
}