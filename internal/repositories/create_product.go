package repositories

import (
	"context"
	"database/sql"
	"fmt"

	"golang-mysql/internal/models"
)

func CreateProduct(ctx context.Context, dbConn *sql.DB, product models.Product) (int, error) {
	query := `
		INSERT INTO products (name, description, price, quantity, category_id, image, unit, status)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id
	`

	var id int
	err := dbConn.QueryRowContext(ctx, query,
		product.Name,
		product.Description,
		product.Price,
		product.Quantity,
		product.CategoryID,
		product.Image,
		product.Unit,
		product.Status,
	).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("gagal menambahkan produk: %w", err)
	}

	return id, nil
}
