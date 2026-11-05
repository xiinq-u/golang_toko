package repositories

import (
	"context"
	"database/sql"
	"fmt"
)

func CreateCategory(ctx context.Context, dbConn *sql.DB, name string) (int, error) {
	const query = `INSERT INTO categories (name) VALUES ($1) RETURNING id`

	var id int
	if err := dbConn.QueryRowContext(ctx, query, name).Scan(&id); err != nil {
		return 0, fmt.Errorf("gagal menambahkan kategori: %w", err)
	}

	return id, nil
}
