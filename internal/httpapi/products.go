package httpapi

import (
	"context"
	"crypto/subtle"
	"database/sql"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"golang-mysql/internal/models"
	"golang-mysql/internal/repositories"
)

type Handler struct {
	db            *sql.DB
	adminToken    string
	allowedOrigin string
}

type createProductRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Quantity    int    `json:"quantity"`
	CategoryID  *int   `json:"category_id"`
	Image       string `json:"image"`
	Unit        string `json:"unit"`
	Status      string `json:"status"`
}

func New(dbConn *sql.DB, adminToken, allowedOrigin string) http.Handler {
	return &Handler{db: dbConn, adminToken: adminToken, allowedOrigin: allowedOrigin}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if origin := r.Header.Get("Origin"); origin != "" && origin == h.allowedOrigin {
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
		w.Header().Add("Vary", "Origin")
	}

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	if r.URL.Path != "/products" {
		http.NotFound(w, r)
		return
	}

	switch r.Method {
	case http.MethodGet:
		h.listProducts(w, r)
	case http.MethodPost:
		if !h.authorized(r) {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}
		h.createProduct(w, r)
	default:
		w.Header().Set("Allow", "GET, POST, OPTIONS")
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *Handler) authorized(r *http.Request) bool {
	if h.adminToken == "" {
		return false
	}
	value := r.Header.Get("Authorization")
	token := strings.TrimPrefix(value, "Bearer ")
	return token != value && subtle.ConstantTimeCompare([]byte(token), []byte(h.adminToken)) == 1
}

func (h *Handler) listProducts(w http.ResponseWriter, r *http.Request) {
	products, err := repositories.ListProducts(h.db)
	if err != nil {
		log.Printf("gagal mengambil produk: %v", err)
		http.Error(w, "gagal mengambil produk", http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusOK, products)
}

func (h *Handler) createProduct(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, 1<<20)
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	var request createProductRequest
	if err := decoder.Decode(&request); err != nil {
		http.Error(w, "body JSON tidak valid", http.StatusBadRequest)
		return
	}
	if err := decoder.Decode(&struct{}{}); !errors.Is(err, io.EOF) {
		http.Error(w, "body harus berisi satu objek JSON", http.StatusBadRequest)
		return
	}

	request.Name = strings.TrimSpace(request.Name)
	request.Unit = strings.TrimSpace(request.Unit)
	request.Status = strings.TrimSpace(request.Status)
	if request.Name == "" || request.Unit == "" || request.Status == "" || request.Price < 0 || request.Quantity < 0 {
		http.Error(w, "name, unit, dan status wajib diisi; price dan quantity tidak boleh negatif", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	product := models.Product{
		Name:        request.Name,
		Description: strings.TrimSpace(request.Description),
		Price:       request.Price,
		Quantity:    request.Quantity,
		CategoryID:  request.CategoryID,
		Image:       strings.TrimSpace(request.Image),
		Unit:        request.Unit,
		Status:      request.Status,
	}
	id, err := repositories.CreateProduct(ctx, h.db, product)
	if err != nil {
		log.Printf("gagal membuat produk: %v", err)
		http.Error(w, "gagal menyimpan produk", http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusCreated, map[string]any{
		"id":          id,
		"name":        product.Name,
		"description": product.Description,
		"price":       product.Price,
		"quantity":    product.Quantity,
		"category_id": product.CategoryID,
		"image":       product.Image,
		"unit":        product.Unit,
		"status":      product.Status,
	})
}

func writeJSON(w http.ResponseWriter, status int, value any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(value); err != nil {
		log.Printf("gagal menulis response JSON: %v", err)
	}
}
