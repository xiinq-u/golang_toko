package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func Open(dsn string) (*sql.DB, error) {
	return sql.Open("postgres", dsn)
}

func Ping(dbConn *sql.DB) error {
	return dbConn.Ping()
}
