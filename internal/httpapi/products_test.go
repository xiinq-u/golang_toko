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
