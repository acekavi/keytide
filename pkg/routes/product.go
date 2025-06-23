package routes

import (
	"github.com/acekavi/keytide/internal/handlers"
	"github.com/gorilla/mux"
)

// RegisterProductRoutes registers all product-related routes
func RegisterProductRoutes(router *mux.Router, productHandler *handlers.ProductHandler) {
    // Create a subrouter for products
    productRouter := router.PathPrefix("/products").Subrouter()

    // Register routes on the subrouter
    productRouter.HandleFunc("", productHandler.GetProducts).Methods("GET")
    productRouter.HandleFunc("/{id}", productHandler.GetProduct).Methods("GET")
    productRouter.HandleFunc("", productHandler.CreateProduct).Methods("POST")
    productRouter.HandleFunc("/{id}", productHandler.UpdateProduct).Methods("PUT")
    productRouter.HandleFunc("/{id}", productHandler.DeleteProduct).Methods("DELETE")
}
