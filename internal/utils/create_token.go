package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenType int

const (
	ACCESS_TOKEN TokenType = iota
	REFRESH_TOKEN
)

type TokenPayload struct {
	Exp          time.Duration
	Email        string
	Role         string
	TokenVersion int
}

func CreateToken(payload TokenPayload, tokenType TokenType) (string, error) {
	secret := os.Getenv("REFRESH_SECRET_KEY")
	if tokenType != REFRESH_TOKEN {
		secret = os.Getenv("ACCESS_SECRET_KEY")
	}

	claims := jwt.MapClaims{
		"exp": time.Now().Add(payload.Exp).Unix(),
	}

	if tokenType == ACCESS_TOKEN {
		claims["sub"] = payload.Email
		claims["role"] = payload.Role
		claims["version"] = payload.TokenVersion
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))

	return tokenString, err
}
