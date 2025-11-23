package middlewares

import (
	"context"
	"net/http"
	"strings"

	"github.com/TheGauravsahu/school-api/internal/utils"
)

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			utils.WriteError(w, http.StatusUnauthorized, "missing or invalid authorization header")
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := utils.VerifyToken(tokenString)
		if err != nil {
			utils.WriteError(w, http.StatusUnauthorized, "invalid token")
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, "user", claims)
		next(w, r.WithContext(ctx))
	}
}
