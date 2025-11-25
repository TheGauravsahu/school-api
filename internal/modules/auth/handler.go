package auth

import (
	"net/http"

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
