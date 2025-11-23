package utils

import (
	"context"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

func SetUserContext(ctx context.Context, claims jwt.MapClaims) context.Context {
	return context.WithValue(ctx, "user", claims)
}

func GetUserContext(r *http.Request) jwt.MapClaims {
	val := r.Context().Value("user")
	if claims, ok := val.(jwt.MapClaims); ok {
		return claims
	}
	return nil
}
