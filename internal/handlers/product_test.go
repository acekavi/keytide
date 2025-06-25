package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/acekavi/keytide/internal/models"
)

// MockProductRepository is a mock implementation of the ProductRepository interface
type MockProductRepository struct {
	products []models.Product
	err      error
}

func (m *MockProductRepository) GetAll() ([]models.Product, error) {
	if m.err != nil {
		return nil, m.err
	}
	return m.products, nil
}

func (m *MockProductRepository) GetByID(id string) (models.Product, error) {
	if m.err != nil {
		return models.Product{}, m.err
	}
	for _, p := range m.products {
		if p.ID == id {
			return p, nil
		}
	}
	return models.Product{}, errors.New("product not found")
}

func (m *MockProductRepository) Create(product models.Product) error {
	if m.err != nil {
		return m.err
	}
	m.products = append(m.products, product)
	return nil
}

func (m *MockProductRepository) Update(product models.Product) error {
	if m.err != nil {
		return m.err
	}
	for i, p := range m.products {
		if p.ID == product.ID {
			m.products[i] = product
			return nil
		}
	}
	return errors.New("product not found")
}

func (m *MockProductRepository) Delete(id string) error {
	if m.err != nil {
		return m.err
	}
	for i, p := range m.products {
		if p.ID == id {
			m.products = append(m.products[:i], m.products[i+1:]...)
			return nil
		}
	}
	return errors.New("product not found")
}

func TestGetProducts(t *testing.T) {
	// Test cases
	testCases := []struct {
		name           string
		mockProducts   []models.Product
		mockError      error
		expectedStatus int
		expectedLen    int
	}{
		{
			name: "success",
			mockProducts: []models.Product{
				{ID: "1", Name: "Laptop", Price: 999.99},
				{ID: "2", Name: "Headphones", Price: 99.99},
			},
			mockError:      nil,
			expectedStatus: http.StatusOK,
			expectedLen:    2,
		},
		{
			name:           "empty products",
			mockProducts:   []models.Product{},
			mockError:      nil,
			expectedStatus: http.StatusOK,
			expectedLen:    0,
		},
		{
			name:           "repository error",
			mockProducts:   nil,
			mockError:      errors.New("database error"),
			expectedStatus: http.StatusInternalServerError,
			expectedLen:    0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create mock repository
			mockRepo := &MockProductRepository{
				products: tc.mockProducts,
				err:      tc.mockError,
			}

			// Create handler with mock repository
			handler := NewProductHandler(mockRepo)

			// Create request
			req, err := http.NewRequest("GET", "/products", nil)
			if err != nil {
				t.Fatal(err)
			}

			// Create response recorder
			rr := httptest.NewRecorder()

			// Call handler
			handler.GetProducts(rr, req)

			// Check status code
			if status := rr.Code; status != tc.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tc.expectedStatus)
			}

			// If expecting success response, check the content
			switch tc.expectedStatus {
			case http.StatusOK:
				// Check content type
				expectedContentType := "application/json"
				if contentType := rr.Header().Get("Content-Type"); contentType != expectedContentType {
					t.Errorf("handler returned wrong content type: got %v want %v",
						contentType, expectedContentType)
				}

				// Unmarshal and check response
				var products []models.Product
				if err := json.Unmarshal(rr.Body.Bytes(), &products); err != nil {
					t.Errorf("Failed to unmarshal response: %v", err)
				}

				// Check length matches expected
				if len(products) != tc.expectedLen {
					t.Errorf("Expected %d products, got %d", tc.expectedLen, len(products))
				}

				// Verify products match expected data
				if tc.expectedLen > 0 {
					for i, expectedProduct := range tc.mockProducts {
						if products[i].ID != expectedProduct.ID ||
							products[i].Name != expectedProduct.Name ||
							products[i].Price != expectedProduct.Price {
							t.Errorf("Product %d does not match expected data", i)
						}
					}
				}
			case http.StatusInternalServerError:
				// For error responses, check error message
				var errorResp map[string]string
				if err := json.Unmarshal(rr.Body.Bytes(), &errorResp); err != nil {
					t.Errorf("Failed to unmarshal error response: %v", err)
				}

				if _, exists := errorResp["error"]; !exists {
					t.Errorf("Expected error field in response")
				}
			}
		})
	}
}
