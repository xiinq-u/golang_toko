package repositories

import (
	"database/sql"
	"fmt"

	"golang-mysql/internal/models"
)

func ListCategories(dbConn *sql.DB) ([]models.Category, error) {
	query := `
		SELECT id, name, description, image, status, created_at, updated_at
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
		var description, image, status sql.NullString
		if err := rows.Scan(&c.ID, &c.Name, &description, &image, &status, &c.CreatedAt, &c.UpdatedAt); err != nil {
			return nil, fmt.Errorf("gagal scan category: %w", err)
		}
		if description.Valid {
			c.Description = description.String
		}
		if image.Valid {
			c.Image = image.String
		}
		if status.Valid {
			c.Status = status.String
		}
		items = append(items, c)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterasi categories: %w", err)
	}

	return items, nil
}
