package repositories

import (
	"database/sql"
	"fmt"

	"golang-mysql/internal/models"
)

func ListTransactionItems(dbConn *sql.DB) ([]models.TransactionItem, error) {
	query := `
		SELECT id, transaction_id, product_id, quantity, price, created_at
		FROM transaction_items
		ORDER BY id
	`

	rows, err := dbConn.Query(query)
	if err != nil {
		return nil, fmt.Errorf("gagal membaca transaction_items: %w", err)
	}
	defer rows.Close()

	items := []models.TransactionItem{}
	for rows.Next() {
		var ti models.TransactionItem
		if err := rows.Scan(&ti.ID, &ti.TransactionID, &ti.ProductID, &ti.Quantity, &ti.Price, &ti.CreatedAt); err != nil {
			return nil, fmt.Errorf("gagal scan transaction_item: %w", err)
		}
		items = append(items, ti)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterasi transaction_items: %w", err)
	}

	return items, nil
}
