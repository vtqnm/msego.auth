package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/vtqnm/msego.auth/internal/config"
)

func GenerateToken() (string, error) {
	claims := jwt.MapClaims{
		"exp": time.Now().Add(config.GetInstance().TokenTTL).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	tokenString, err := token.SignedString([]byte("Secret key"))
	if err != nil {
		return "", err
	}

	return tokenString, err
}
