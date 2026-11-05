package repositories

import (
	"database/sql"
	"fmt"

	"golang-mysql/internal/models"
)

func ListTransactions(dbConn *sql.DB) ([]models.Transaction, error) {
	query := `
		SELECT id, user_id, total, status, created_at, updated_at
		FROM transactions
		ORDER BY id
	`

	rows, err := dbConn.Query(query)
	if err != nil {
		return nil, fmt.Errorf("gagal membaca transactions: %w", err)
	}
	defer rows.Close()

	items := []models.Transaction{}
	for rows.Next() {
		var t models.Transaction
		if err := rows.Scan(&t.ID, &t.UserID, &t.Total, &t.Status, &t.CreatedAt, &t.UpdatedAt); err != nil {
			return nil, fmt.Errorf("gagal scan transaction: %w", err)
		}
		items = append(items, t)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterasi transactions: %w", err)
	}

	return items, nil
}
