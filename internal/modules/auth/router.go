package auth

import (
	"net/http"
)

func Router(handler *Handler) {
	http.HandleFunc("POST /api/auth/register-admin", handler.RegisterAdmin)
	http.HandleFunc("POST /api/auth/refresh", handler.Refresh)
	http.HandleFunc("POST /api/auth/login", handler.Login)

}
