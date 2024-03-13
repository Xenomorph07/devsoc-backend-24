package middleware

import (
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func Protected() echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("ACCESS_SECRET_KEY")),
		TokenLookup: "cookie:" + os.Getenv("ACCESS_COOKIE_NAME"),
	})
}

func Refresh() echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("REFRESH_SECRET_KEY")),
		TokenLookup: "cookie:" + os.Getenv("REFRESH_COOKIE_NAME"),
	})
}
