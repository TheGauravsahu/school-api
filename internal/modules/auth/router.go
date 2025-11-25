package auth

import "net/http"

func Router(handler *Handler) {
	http.HandleFunc("POST /api/auth/register-admin", handler.RegisterAdmin)
	http.HandleFunc("POST /api/auth/login", handler.Login)

}
