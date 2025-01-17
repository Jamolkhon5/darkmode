package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Jamolkhon5/darkmode/internal/auth"
	"github.com/Jamolkhon5/darkmode/internal/models"
	"github.com/Jamolkhon5/darkmode/internal/repository"
)

type Handler struct {
	themeRepo repository.Theme
}

func NewHandler(repo repository.Theme) *Handler {
	return &Handler{themeRepo: repo}
}

func (h *Handler) UpdateTheme(w http.ResponseWriter, r *http.Request) {
	userID, err := auth.VerifyToken(r)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var req models.ThemeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.Theme != "light" && req.Theme != "dark-only" {
		http.Error(w, "Invalid theme value", http.StatusBadRequest)
		return
	}

	theme, err := h.themeRepo.UpdateTheme(userID, req.Theme)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(theme)
}
