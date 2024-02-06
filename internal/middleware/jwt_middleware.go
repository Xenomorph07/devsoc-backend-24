package middleware

import (
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func Protected() echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		SigningKey:  os.Getenv("ACCESS_SECRET_KEY"),
		TokenLookup: "header:Authorization:Bearer ,cookie:access_token",
	})
}
