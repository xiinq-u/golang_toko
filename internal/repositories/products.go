package repositories

import (
	"database/sql"
	"fmt"

	"golang-mysql/internal/models"
)

func ListProducts(dbConn *sql.DB) ([]models.Product, error) {
	query := `
		SELECT p.id, p.name, p.description, p.price, p.quantity, p.category_id,
		       c.name, p.image, p.unit, p.status, p.created_at, p.updated_at
		FROM products p
		LEFT JOIN categories c ON c.id = p.category_id
		ORDER BY p.id
	`

	rows, err := dbConn.Query(query)
	if err != nil {
		return nil, fmt.Errorf("gagal membaca products: %w", err)
	}
	defer rows.Close()

	items := []models.Product{}
	for rows.Next() {
		var p models.Product
		var description, categoryName, image, unit, status sql.NullString
		var categoryID sql.NullInt64
		var createdAt, updatedAt sql.NullTime
		if err := rows.Scan(
			&p.ID,
			&p.Name,
			&description,
			&p.Price,
			&p.Quantity,
			&categoryID,
			&categoryName,
			&image,
			&unit,
			&status,
			&createdAt,
			&updatedAt,
		); err != nil {
			return nil, fmt.Errorf("gagal scan product: %w", err)
		}
		if description.Valid {
			p.Description = description.String
		}
		if categoryName.Valid {
			p.CategoryName = categoryName.String
		}
		if categoryID.Valid {
			id := int(categoryID.Int64)
			p.CategoryID = &id
		}
		if image.Valid {
			p.Image = image.String
		}
		if unit.Valid {
			p.Unit = unit.String
		}
		if status.Valid {
			p.Status = status.String
		}
		if createdAt.Valid {
			p.CreatedAt = createdAt.Time
		}
		if updatedAt.Valid {
			p.UpdatedAt = updatedAt.Time
		}
		items = append(items, p)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterasi products: %w", err)
	}

	return items, nil
}
