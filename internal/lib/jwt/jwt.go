package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(duration time.Duration) (string, error) {
	claims := jwt.MapClaims{
		"exp": time.Now().Add(duration).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	tokenString, err := token.SignedString([]byte("Secret key"))
	if err != nil {
		return "", err
	}

	return tokenString, err
}
