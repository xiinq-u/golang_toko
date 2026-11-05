package main

import (
	"log"
	"net/http"
	"time"

	"golang-mysql/internal/config"
	"golang-mysql/internal/db"
	"golang-mysql/internal/httpapi"
)

func main() {
	dsn, err := config.GetDSN()
	if err != nil {
		log.Fatal(err)
	}

	dbConn, err := db.Open(dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer dbConn.Close()

	if err := db.Ping(dbConn); err != nil {
		log.Fatalf("Gagal koneksi ke Supabase: %v", err)
	}
	adminToken, err := config.GetAdminAPIToken()
	if err != nil {
		log.Fatal(err)
	}

	server := &http.Server{
		Addr:              ":" + config.GetPort(),
		Handler:           httpapi.New(dbConn, adminToken, config.GetAllowedOrigin()),
		ReadHeaderTimeout: 5 * time.Second,
		IdleTimeout:       60 * time.Second,
	}
	log.Printf("API produk berjalan di port %s", config.GetPort())
	log.Fatal(server.ListenAndServe())
}
