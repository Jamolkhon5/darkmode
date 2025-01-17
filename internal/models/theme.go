package models

type Theme struct {
	ID        int    `json:"-" db:"id"`
	UserID    string `json:"userId" db:"user_id"`
	Theme     string `json:"theme" db:"theme"`
	UpdatedAt string `json:"updatedAt" db:"updated_at"`
}

type ThemeRequest struct {
	Theme string `json:"theme"`
}
