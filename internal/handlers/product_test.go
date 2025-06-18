package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetProducts(t *testing.T) {
    // Create a request
    req, err := http.NewRequest("GET", "/products", nil)
    if err != nil {
        t.Fatal(err)
    }

    // Create a ResponseRecorder to record the response
    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(GetProducts)

    // Call the handler
    handler.ServeHTTP(rr, req)

    // Check status code
    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusOK)
    }

    // Check content type
    expectedContentType := "application/json"
    if contentType := rr.Header().Get("Content-Type"); contentType != expectedContentType {
        t.Errorf("handler returned wrong content type: got %v want %v",
            contentType, expectedContentType)
    }

    // Unmarshal and check response
    var products []map[string]any
    if err := json.Unmarshal(rr.Body.Bytes(), &products); err != nil {
        t.Errorf("Failed to unmarshal response: %v", err)
    }
    
    // We expect some products in the initial response
    if len(products) == 0 {
        t.Errorf("Expected non-empty products array")
    }
}