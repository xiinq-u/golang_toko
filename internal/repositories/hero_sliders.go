package repositories

import (
	"database/sql"
	"fmt"

	"golang-mysql/internal/models"
)

func ListHeroSliders(dbConn *sql.DB) ([]models.HeroSlider, error) {
	query := `
		SELECT id, title, description, image, created_at, updated_at
		FROM hero_sliders
		ORDER BY id
	`

	rows, err := dbConn.Query(query)
	if err != nil {
		return nil, fmt.Errorf("gagal membaca hero_sliders: %w", err)
	}
	defer rows.Close()

	items := []models.HeroSlider{}
	for rows.Next() {
		var hs models.HeroSlider
		var description, image sql.NullString
		if err := rows.Scan(&hs.ID, &hs.Title, &description, &image, &hs.CreatedAt, &hs.UpdatedAt); err != nil {
			return nil, fmt.Errorf("gagal scan hero_slider: %w", err)
		}
		if description.Valid {
			hs.Description = description.String
		}
		if image.Valid {
			hs.Image = image.String
		}
		items = append(items, hs)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterasi hero_sliders: %w", err)
	}

	return items, nil
}
