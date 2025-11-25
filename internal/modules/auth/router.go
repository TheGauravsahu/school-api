package auth

import "net/http"

func Router(handler *Handler) {
	http.HandleFunc("/api/auth/register-admin", handler.RegisterAdmin)
}
