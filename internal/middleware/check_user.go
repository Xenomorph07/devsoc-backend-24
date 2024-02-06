package middleware

import (
	"log/slog"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func AuthUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token, ok := c.Get("user").(*jwt.Token)

		if !ok {
			return c.JSON(http.StatusNotAcceptable, map[string]string{
				"message": "JWT is invalid or missing",
				"status":  "validating jwt",
			})
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"message": "malformed jwt",
				"status":  "checking claims",
			})
		}

		slog.Info("The claims for the passed in JWT are: ")

		for key, value := range claims {
			slog.Info(key + ":" + value.(string))
		}

		return next(c)
	}
}
