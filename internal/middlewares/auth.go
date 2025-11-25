package middlewares

import (
	"context"
	"net/http"
	"strings"

	"github.com/TheGauravsahu/school-api/internal/utils"
)

func AuthMiddleware(next http.HandlerFunc, allowedRoles ...string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			utils.WriteError(w, http.StatusUnauthorized, "missing or invalid authorization header")
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := utils.VerifyAcessToken(tokenString)
		role := claims["role"].(string)
		if err != nil {
			utils.WriteError(w, http.StatusUnauthorized, "invalid token")
			return
		}

		if len(allowedRoles) > 0 {
			allowed := false
			for _, r := range allowedRoles {
				if r == role {
					allowed = true
					break
				}
			}
			if !allowed {
				utils.WriteError(w, http.StatusForbidden, "forbidden: insufficient permissions")
				return
			}
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, "user", claims)
		next(w, r.WithContext(ctx))
	}
}
