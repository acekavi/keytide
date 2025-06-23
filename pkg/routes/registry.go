package routes

import (
	"github.com/acekavi/keytide/internal/handlers"
	"github.com/gorilla/mux"
)

// RegisterRoutes sets up all API routes
func RegisterRoutes(router *mux.Router, productHandler *handlers.ProductHandler) {
    // Create API v1 subrouter
    apiV1 := router.PathPrefix("/v1").Subrouter()

    // Register resource routes
    RegisterProductRoutes(apiV1, productHandler)

    // You can add more resource routes here as your API grows
    // RegisterUserRoutes(apiV1, userHandler)
    // RegisterAuthRoutes(apiV1, authHandler)
}
