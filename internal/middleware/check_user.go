package middleware

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"

	"github.com/CodeChefVIT/devsoc-backend-24/internal/database"
	services "github.com/CodeChefVIT/devsoc-backend-24/internal/services/user"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
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

		email := claims["sub"].(string)
		tokenVersionStr, err := database.RedisClient.Get(email)
		if err != nil {
			if err == redis.Nil {
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"message": "token expired",
					"status":  "failure",
				})
			}
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"message": err.Error(),
				"status":  "get from redis",
			})
		}

		tokenVersion, _ := strconv.Atoi(tokenVersionStr)

		if int(claims["version"].(float64)) != tokenVersion {
			return c.JSON(http.StatusForbidden, map[string]string{
				"messsage": "invalid token",
				"status":   "failure",
			})
		}

		user, err := services.FindUserByEmail(email)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return c.JSON(http.StatusNotFound, map[string]string{
					"message": "user does not exist",
					"status":  "failure",
				})
			}
		}

		c.Set("user", user)

		return next(c)
	}
}
