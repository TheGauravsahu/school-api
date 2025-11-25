package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var AccessSecret = []byte("ACCESS_SECRET_32_CHARS")
var RefreshSecret = []byte("REFRESH_SECRET_32_CHARS")

func GenerateAccessToken(userID uint, schoolID uint, username string, role string) (string, error) {
	claims := jwt.MapClaims{
		"user_id":   userID,
		"school_id": schoolID,
		"username":  username,
		"role":      role,
		"exp":       time.Now().Add(15 * time.Minute).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(AccessSecret)
}

func GenerateRefreshToken(userID uint, schoolID uint, username string, role string) (string, error) {
	claims := jwt.MapClaims{
		"user_id":   userID,
		"school_id": schoolID,
		"username":  username,
		"role":      role,
		"exp":       time.Now().Add(7 * 24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(RefreshSecret)
}

func VerifyRefreshToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (any, error) {
		if t.Method != jwt.SigningMethodHS256 {
			return nil, errors.New("invalid signing method")
		}
		return RefreshSecret, nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("invalid token")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid token claims")
	}
	return claims, nil
}

func VerifyAcessToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (any, error) {
		if t.Method != jwt.SigningMethodHS256 {
			return nil, errors.New("invalid signing method")
		}
		return AccessSecret, nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid token claims")
	}
	return claims, nil
}
