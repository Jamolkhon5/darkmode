package repository

import "github.com/Jamolkhon5/darkmode/internal/models"

type Repository interface {
	UpdateTheme(userID string, theme string) (*models.Theme, error)
}
