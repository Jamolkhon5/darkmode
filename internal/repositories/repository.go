package repository

import (
	"github.com/jmoiron/sqlx"
	"your_username/darkmode/internal/models"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) UpdateTheme(userID string, theme string) (*models.Theme, error) {
	query := `
        INSERT INTO themes (user_id, theme) 
        VALUES ($1, $2)
        ON CONFLICT (user_id) 
        DO UPDATE SET theme = EXCLUDED.theme, updated_at = CURRENT_TIMESTAMP
        RETURNING *`

	var updatedTheme models.Theme
	err := r.db.Get(&updatedTheme, query, userID, theme)
	if err != nil {
		return nil, err
	}

	return &updatedTheme, nil
}
