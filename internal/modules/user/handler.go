package user

import (
	"net/http"

	"github.com/TheGauravsahu/school-api/internal/utils"
)

type Handler struct {
	repo *Repository
}

func NewHandler(repo *Repository) *Handler {
	return &Handler{repo: repo}
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var userBody User

	if ok := utils.ParseJson(w, r, &userBody); !ok {
		return
	}

	user, _ := h.repo.GetUserByUsername(userBody.Username)
	if user != nil {
		utils.WriteError(w, http.StatusNotFound, "user already exists.")
		return
	}

	hashed, err := utils.HashPassword(userBody.Password)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, "Internal server error")
		return
	}
	userBody.Password = hashed

	err = h.repo.CreateUser(&userBody)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, "Failed to create user")
		return
	}

	utils.WriteJson(w, http.StatusCreated, map[string]string{
		"message": "user created successfully",
	})
}
