package user

import (
	"net/http"

	"github.com/TheGauravsahu/school-api/internal/middlewares"
)

func Router(h *Handler) {
	http.HandleFunc("/api/users/register", h.CreateUser)
	http.HandleFunc("/api/users/login", h.LoginUser)
	http.HandleFunc("/api/users/me", middlewares.AuthMiddleware(h.GetProfile))
}
