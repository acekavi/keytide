package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/acekavi/keytide/internal/models"
	"github.com/acekavi/keytide/internal/repository"
	"github.com/acekavi/keytide/pkg/logger"
	"github.com/acekavi/keytide/pkg/utils"
	"github.com/acekavi/keytide/pkg/validator"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

// ProductHandler handles product-related requests
type ProductHandler struct {
    repo repository.ProductRepository
}

// NewProductHandler creates a new product handler
func NewProductHandler(repo repository.ProductRepository) *ProductHandler {
    return &ProductHandler{
        repo: repo,
    }
}

// ProductRequest represents a product request body
type ProductRequest struct {
    Name  string  `json:"name" validate:"required,min=3,max=100"`
    Price float64 `json:"price" validate:"required,min=0.01"`
}

// GetProducts returns all products
func (h *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
    products, err := h.repo.GetAll()
    if err != nil {
        logger.Error("Failed to get products", zap.Error(err))
        utils.JSONError(w, http.StatusInternalServerError, "Failed to retrieve products")
        return
    }

    utils.JSONResponse(w, http.StatusOK, products)
}

// GetProduct returns a specific product
func (h *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
    id := mux.Vars(r)["id"]
    
    product, err := h.repo.GetByID(id)
    if err != nil {
        if strings.Contains(err.Error(), "not found") {
            utils.JSONError(w, http.StatusNotFound, "Product not found")
            return
        }
        logger.Error("Failed to get product", zap.Error(err), zap.String("id", id))
        utils.JSONError(w, http.StatusInternalServerError, "Failed to retrieve product")
        return
    }

    utils.JSONResponse(w, http.StatusOK, product)
}

// CreateProduct creates a new product
func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
    var req ProductRequest
    
    // Decode request
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        utils.JSONError(w, http.StatusBadRequest, "Invalid request body")
        return
    }
    
    // Validate request
    if errors := validator.Validate(req); len(errors) > 0 {
        utils.JSONResponse(w, http.StatusBadRequest, map[string]interface{}{
            "errors": errors,
        })
        return
    }
    
    // Create product
    product := models.Product{
        ID:    utils.GenerateID(),
        Name:  req.Name,
        Price: req.Price,
    }
    
    if err := h.repo.Create(product); err != nil {
        logger.Error("Failed to create product", zap.Error(err))
        utils.JSONError(w, http.StatusInternalServerError, "Failed to create product")
        return
    }
    
    utils.JSONResponse(w, http.StatusCreated, product)
}

// UpdateProduct updates an existing product
func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
    id := mux.Vars(r)["id"]
    
    var req ProductRequest
    
    // Decode request
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        utils.JSONError(w, http.StatusBadRequest, "Invalid request body")
        return
    }
    
    // Validate request
    if errors := validator.Validate(req); len(errors) > 0 {
        utils.JSONResponse(w, http.StatusBadRequest, map[string]interface{}{
            "errors": errors,
        })
        return
    }
    
    // Update product
    product := models.Product{
        ID:    id,
        Name:  req.Name,
        Price: req.Price,
    }
    
    if err := h.repo.Update(product); err != nil {
        if strings.Contains(err.Error(), "not found") {
            utils.JSONError(w, http.StatusNotFound, "Product not found")
            return
        }
        logger.Error("Failed to update product", zap.Error(err), zap.String("id", id))
        utils.JSONError(w, http.StatusInternalServerError, "Failed to update product")
        return
    }
    
    utils.JSONResponse(w, http.StatusOK, product)
}

// DeleteProduct deletes a product
func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
    id := mux.Vars(r)["id"]
    
    if err := h.repo.Delete(id); err != nil {
        if strings.Contains(err.Error(), "not found") {
            utils.JSONError(w, http.StatusNotFound, "Product not found")
            return
        }
        logger.Error("Failed to delete product", zap.Error(err), zap.String("id", id))
        utils.JSONError(w, http.StatusInternalServerError, "Failed to delete product")
        return
    }
    
    utils.JSONResponse(w, http.StatusNoContent, nil)
}