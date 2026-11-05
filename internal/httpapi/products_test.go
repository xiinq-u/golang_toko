package httpapi

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCreateProductRequiresAdminToken(t *testing.T) {
	handler := New(nil, "expected-token", "http://localhost:3000")
	request := httptest.NewRequest(http.MethodPost, "/products", strings.NewReader(`{"name":"Test","price":1,"quantity":1,"unit":"pcs","status":"active"}`))
	response := httptest.NewRecorder()

	handler.ServeHTTP(response, request)

	if response.Code != http.StatusUnauthorized {
		t.Fatalf("expected status %d, got %d", http.StatusUnauthorized, response.Code)
	}
}

func TestRootListsVersionedRoutes(t *testing.T) {
	handler := New(nil, "expected-token", "http://localhost:3000")
	request := httptest.NewRequest(http.MethodGet, "/", nil)
	response := httptest.NewRecorder()

	handler.ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, response.Code)
	}
	if !strings.Contains(response.Body.String(), "/api/v1/categories") {
		t.Fatalf("expected route list in response, got %s", response.Body.String())
	}
}

func TestTransactionsRequiresAdminToken(t *testing.T) {
	handler := New(nil, "expected-token", "http://localhost:3000")
	request := httptest.NewRequest(http.MethodGet, "/api/v1/transactions", nil)
	response := httptest.NewRecorder()

	handler.ServeHTTP(response, request)

	if response.Code != http.StatusUnauthorized {
		t.Fatalf("expected status %d, got %d", http.StatusUnauthorized, response.Code)
	}
}
