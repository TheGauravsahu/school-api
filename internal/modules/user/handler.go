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

func (h *Handler) LoginUser(w http.ResponseWriter, r *http.Request) {
	var loginBody User
	if ok := utils.ParseJson(w, r, &loginBody); !ok {
		return
	}

	user, err := h.repo.GetUserByUsername(loginBody.Username)
	if err != nil || user == nil {
		utils.WriteError(w, http.StatusNotFound, "user not found.")
		return
	}

	// check pass
	if !utils.ComparePassword(user.Password, loginBody.Password) {
		utils.WriteError(w, http.StatusUnauthorized, "invalid credentials")
		return
	}

	token, err := utils.GenerateToken(user.Username, user.Role)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, "failed to login")
		return
	}

	user.Password = ""

	utils.WriteJson(w, http.StatusOK, map[string]interface{}{
		"message": "User logged in",
		"user":    user,
		"token":   token,
	})
}

func (h *Handler) GetProfile(w http.ResponseWriter, r *http.Request) {
	claims := utils.GetUserContext(r)
	if claims == nil {
		utils.WriteError(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	utils.WriteJson(w, http.StatusOK, map[string]interface{}{
		"username": claims["username"],
		"role":     claims["role"],
	})
}
