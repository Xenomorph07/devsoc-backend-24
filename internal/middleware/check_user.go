package middleware

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"

	"github.com/CodeChefVIT/devsoc-backend-24/internal/database"
	services "github.com/CodeChefVIT/devsoc-backend-24/internal/services/user"
)

func AuthUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token, ok := c.Get("user").(*jwt.Token)

		if !ok {
			return c.JSON(http.StatusNotAcceptable, map[string]string{
				"message": "JWT is invalid or missing",
				"status":  "fail",
			})
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"message": "malformed jwt",
				"status":  "fail",
				"data": map[string]bool{
					"malformed": true,
				},
			})
		}

		email := claims["sub"].(string)

		tokenVersionStr, err := database.RedisClient.Get("token_version:" + email)
		if err != nil {
			if err == redis.Nil {
				return c.JSON(http.StatusUnauthorized, map[string]interface{}{
					"message": "token expired",
					"status":  "fail",
					"data": map[string]bool{
						"token_expired": true,
					},
				})
			}

			return c.JSON(http.StatusInternalServerError, map[string]string{
				"message": err.Error(),
				"status":  "error",
			})
		}

		tokenVersion, _ := strconv.Atoi(tokenVersionStr)

		if int(claims["version"].(float64)) != tokenVersion {
			return c.JSON(http.StatusForbidden, map[string]interface{}{
				"messsage": "invalid token",
				"status":   "fail",
				"data": map[string]interface{}{
					"invalid_token": true,
				},
			})
		}

		user, err := services.FindUserByEmail(email)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return c.JSON(http.StatusNotFound, map[string]string{
					"message": "user does not exist",
					"status":  "fail",
				})
			}
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"message": err.Error(),
				"status":  "fail",
			})
		}

		if user.IsBanned {
			return c.JSON(http.StatusFailedDependency, map[string]string{
				"message": "user is banned",
				"status":  "fail",
			})
		}

		if !user.IsVerified {
			return c.JSON(http.StatusForbidden, map[string]string{
				"message": "not verified",
				"status":  "fail",
			})
		}

		c.Set("user", &user.User)

		return next(c)
	}
}

func CheckAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token, ok := c.Get("user").(*jwt.Token)

		if !ok {
			return c.JSON(http.StatusNotAcceptable, map[string]string{
				"message": "JWT is invalid or missing",
				"status":  "fail",
			})
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"message": "malformed jwt",
				"status":  "fail",
				"data": map[string]bool{
					"malformed": true,
				},
			})
		}

		email := claims["sub"].(string)

		if claims["role"].(string) != "admin" {
			return c.JSON(http.StatusForbidden, map[string]string{
				"message": "not an admin",
				"status":  "fail",
			})
		}

		tokenVersionStr, err := database.RedisClient.Get("token_version:" + email)
		if err != nil {
			if err == redis.Nil {
				return c.JSON(http.StatusUnauthorized, map[string]interface{}{
					"message": "token expired",
					"status":  "fail",
					"data": map[string]bool{
						"token_expired": true,
					},
				})
			}

			return c.JSON(http.StatusInternalServerError, map[string]string{
				"message": err.Error(),
				"status":  "error",
			})
		}

		tokenVersion, _ := strconv.Atoi(tokenVersionStr)

		if int(claims["version"].(float64)) != tokenVersion {
			return c.JSON(http.StatusForbidden, map[string]interface{}{
				"messsage": "invalid token",
				"status":   "fail",
				"data": map[string]interface{}{
					"invalid_token": true,
				},
			})
		}

		user, err := services.FindUserByEmail(email)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return c.JSON(http.StatusNotFound, map[string]string{
					"message": "user does not exist",
					"status":  "fail",
				})
			}
		}

		if !user.IsVerified {
			return c.JSON(http.StatusForbidden, map[string]string{
				"message": "not verified",
				"status":  "fail",
			})
		}

		c.Set("user", &user.User)

		return next(c)
	}
}
