package config

import (
	"fmt"
	"os"
)

func GetDSN() (string, error) {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		return "", fmt.Errorf("DATABASE_URL belum diatur")
	}
	return dsn, nil
}

func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		return "8080"
	}
	return port
}

func GetAllowedOrigin() string {
	origin := os.Getenv("ALLOWED_ORIGIN")
	if origin == "" {
		return "http://localhost:3000"
	}
	return origin
}

func GetAdminAPIToken() (string, error) {
	token := os.Getenv("ADMIN_API_TOKEN")
	if token == "" {
		return "", fmt.Errorf("ADMIN_API_TOKEN belum diatur")
	}
	return token, nil
}
