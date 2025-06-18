package handlers

import (
	"encoding/json"
	"net/http"

	"keytide/internal/models"
)

// GetProducts returns all products
func GetProducts(w http.ResponseWriter, r *http.Request) {
    // For now, return some hardcoded products
    products := []models.Product{
        {ID: "1", Name: "Laptop", Price: 999.99},
        {ID: "2", Name: "Headphones", Price: 99.99},
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(products)
}