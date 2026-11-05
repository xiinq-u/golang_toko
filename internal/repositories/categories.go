package repositories

import (
	"database/sql"
	"fmt"

	"golang-mysql/internal/models"
)

func ListCategories(dbConn *sql.DB) ([]models.Category, error) {
	query := `
		SELECT id, name, slug, created_at, updated_at
		FROM categories
		ORDER BY id
	`

	rows, err := dbConn.Query(query)
	if err != nil {
		return nil, fmt.Errorf("gagal membaca categories: %w", err)
	}
	defer rows.Close()

	items := []models.Category{}
	for rows.Next() {
		var c models.Category
		if err := rows.Scan(&c.ID, &c.Name, &c.Slug, &c.CreatedAt, &c.UpdatedAt); err != nil {
			return nil, fmt.Errorf("gagal scan category: %w", err)
		}
		items = append(items, c)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterasi categories: %w", err)
	}

	return items, nil
}
