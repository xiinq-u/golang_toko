package repositories

import (
	"database/sql"
	"fmt"

	"golang-mysql/internal/auth"
)

func ListAuthUsers(dbConn *sql.DB) ([]auth.User, error) {
	query := `
		SELECT id
		FROM auth.users
		ORDER BY id
	`

	rows, err := dbConn.Query(query)
	if err != nil {
		return nil, fmt.Errorf("gagal membaca auth.users: %w", err)
	}
	defer rows.Close()

	items := []auth.User{}
	for rows.Next() {
		var u auth.User
		if err := rows.Scan(&u.ID); err != nil {
			return nil, fmt.Errorf("gagal scan auth user: %w", err)
		}
		items = append(items, u)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterasi auth.users: %w", err)
	}

	return items, nil
}
