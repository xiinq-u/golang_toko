package repositories

import (
	"database/sql"
	"fmt"

	"golang-mysql/internal/models"
)

func ListProfiles(dbConn *sql.DB) ([]models.Profile, error) {
	query := `
		SELECT id, auth_user_id, full_name, phone_number, created_at, updated_at
		FROM profiles
		ORDER BY id
	`

	rows, err := dbConn.Query(query)
	if err != nil {
		return nil, fmt.Errorf("gagal membaca profiles: %w", err)
	}
	defer rows.Close()

	items := []models.Profile{}
	for rows.Next() {
		var p models.Profile
		if err := rows.Scan(&p.ID, &p.AuthUserID, &p.FullName, &p.PhoneNumber, &p.CreatedAt, &p.UpdatedAt); err != nil {
			return nil, fmt.Errorf("gagal scan profile: %w", err)
		}
		items = append(items, p)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterasi profiles: %w", err)
	}

	return items, nil
}
