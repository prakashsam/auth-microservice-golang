package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateJWT(userID uint, email string, secret string) (string, error) {
	claims := jwt.MapClaims{
		"sub":   userID,
		"email": email,
		"role":  "user",
		"iss":   "authservice",
		"exp":   time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}
