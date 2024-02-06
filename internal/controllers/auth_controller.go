package controllers

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/CodeChefVIT/devsoc-backend-24/internal/database"
	"github.com/CodeChefVIT/devsoc-backend-24/internal/models"
	services "github.com/CodeChefVIT/devsoc-backend-24/internal/services/user"
	"github.com/CodeChefVIT/devsoc-backend-24/internal/utils"
	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
	"golang.org/x/crypto/bcrypt"
)

func Login(ctx echo.Context) error {
	var payload models.LoginRequest

	if err := ctx.Bind(&payload); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
			"status":  "binding body",
		})
	}

	if err := ctx.Validate(&payload); err != nil {
		return err
	}

	user, err := services.FindUserByEmail(payload.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ctx.JSON(http.StatusNotFound, map[string]string{
				"message": "user does not exist",
				"status":  "failure",
			})
		}
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"messsage": err.Error(),
			"status":   "db error",
		})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password)); err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return ctx.JSON(http.StatusConflict, map[string]string{
				"message": "Invalid password",
				"status":  "failure",
			})
		}
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
			"status":  "bcrypt check",
		})
	}

	tokenVersionStr, err := database.RedisClient.Get(user.Email)
	if err != nil && err != redis.Nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"status":  "redis failure",
			"message": err.Error(),
		})
	}

	tokenVersion, _ := strconv.Atoi(tokenVersionStr)

	accessToken, err := utils.CreateToken(utils.TokenPayload{
		Exp:          time.Minute * 5,
		Email:        user.Email,
		Role:         user.Role,
		TokenVersion: tokenVersion + 1,
	}, utils.ACCESS_TOKEN)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
			"status":  "create token",
		})
	}

	refreshToken, err := utils.CreateToken(utils.TokenPayload{Exp: time.Hour * 1}, utils.REFRESH_TOKEN)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
			"status":  "create token",
		})
	}

	if err := database.RedisClient.Set(user.Email, fmt.Sprint(tokenVersion+1), time.Hour*1); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"status":  "redis failure",
			"message": err.Error(),
		})
	}

	if err := database.RedisClient.Set(refreshToken, user.Email, time.Hour*1); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"status":  "redis failure",
			"message": err.Error(),
		})
	}

	ctx.SetCookie(&http.Cookie{
		Name:     "access_token",
		Value:    accessToken,
		HttpOnly: true,
	})

	ctx.SetCookie(&http.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		HttpOnly: true,
	})

	return ctx.JSON(http.StatusOK, map[string]string{
		"message": "login successful",
		"status":  "success",
	})
}

func Logout(ctx echo.Context) error {
	refreshToken, err := ctx.Cookie("refresh_token")
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
			"status":  "cookie not found",
		})
	}

	email, err := database.RedisClient.Get(refreshToken.Value)
	if err != nil && err != redis.Nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
			"status":  "redis failure",
		})
	}

	if err := database.RedisClient.Delete(refreshToken.Value); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
			"status":  "redis failure",
		})
	}

	if err := database.RedisClient.Delete(email); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
			"status":  "redis failure",
		})
	}

	return ctx.JSON(http.StatusOK, map[string]string{
		"message": "logout successful",
		"status":  "success",
	})
}

func Refresh(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, map[string]string{
		"message": "token refreshed",
		"status":  "success",
	})
}
