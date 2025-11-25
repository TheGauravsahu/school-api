package auth

import (
	"net/http"

	"github.com/TheGauravsahu/school-api/internal/config"
	"github.com/TheGauravsahu/school-api/internal/modules/user"
	"github.com/TheGauravsahu/school-api/internal/utils"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

type RegisterAdminRequest struct {
	SchoolName    string `json:"school_name"`
	SchoolAddress string `json:"school_address"`
	SchoolLogo    string `json:"school_logo"`
	Username      string `json:"username"`
	Password      string `json:"password"`
}

func (h *Handler) RegisterAdmin(w http.ResponseWriter, r *http.Request) {
	var body RegisterAdminRequest

	if ok := utils.ParseJson(w, r, &body); !ok {
		return
	}

	err := h.service.RegisterSchoolAndAdmin(body.SchoolName, body.SchoolAddress, body.SchoolLogo, body.Username, body.Password)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.WriteJson(w, http.StatusCreated, map[string]string{
		"message": "School and admin created",
	})
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if ok := utils.ParseJson(w, r, &body); !ok {
		return
	}

	user, err := h.service.UserRepo.GetUserByUsername(body.Username)
	if err != nil {
		utils.WriteError(w, http.StatusUnauthorized, "Invalid credentials")
		return
	}

	if ok := utils.ComparePassword(user.Password, body.Password); !ok {
		utils.WriteError(w, http.StatusUnauthorized, "Invalid credentials")
		return
	}

	accessToken, _ := utils.GenerateAccessToken(user.ID, user.SchoolID, user.Username, user.Role)
	refreshToken, _ := utils.GenerateRefreshToken(user.ID, user.SchoolID, user.Username, user.Role)

	// save refresh token
	config.DB.Model(&user).Update("refresh_token", refreshToken)

	utils.WriteJson(w, http.StatusOK, map[string]any{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
		"user": map[string]interface{}{
			"id":       user.ID,
			"username": user.Username,
			"role":     user.Role,
		},
	})
}

func (h *Handler) Refresh(w http.ResponseWriter, r *http.Request) {
	var body struct {
		RefreshToken string `json:"refresh_token"`
	}

	if ok := utils.ParseJson(w, r, &body); !ok {
		return
	}

	claims, err := utils.VerifyRefreshToken(body.RefreshToken)
	if err != nil {
		utils.WriteError(w, http.StatusUnauthorized, "Invalid refresh token")
		return
	}

	var user user.User
	userID := uint(claims["user_id"].(float64))
	config.DB.First(&user, userID)

	// cross check DB refresh token
	if user.RefreshToken != body.RefreshToken {
		http.Error(w, "refresh token mismatch", http.StatusUnauthorized)
		return
	}

	accessToken, _ := utils.GenerateAccessToken(user.ID, user.SchoolID, user.Username, user.Role)
	refreshToken, _ := utils.GenerateRefreshToken(user.ID, user.SchoolID, user.Username, user.Role)

	config.DB.Model(&user).Update("refresh_token", refreshToken)

	utils.WriteJson(w, http.StatusOK, map[string]string{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}

func (h *Handler) GetProfile(w http.ResponseWriter, r *http.Request) {
	claims := utils.GetUserContext(r)
	if claims == nil {
		utils.WriteError(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	utils.WriteJson(w, http.StatusOK, map[string]interface{}{
		"user_id":   claims["user_id"],
		"school_id": claims["school_id"],
		"username":  claims["username"],
		"role":      claims["role"],
	})
}
