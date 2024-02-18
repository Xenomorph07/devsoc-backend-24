package controllers

import (
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
	"golang.org/x/crypto/bcrypt"

	"github.com/CodeChefVIT/devsoc-backend-24/internal/database"
	"github.com/CodeChefVIT/devsoc-backend-24/internal/models"
	services "github.com/CodeChefVIT/devsoc-backend-24/internal/services/user"
	"github.com/CodeChefVIT/devsoc-backend-24/internal/utils"
)

func CreateUser(ctx echo.Context) error {
	var payload models.CreateUserRequest

	if err := ctx.Bind(&payload); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
			"status":  "fail",
		})
	}

	if err := ctx.Validate(&payload); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
			"status":  "fail",
		})
	}

	userExists, err := services.CheckUserExists(payload.Email)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
			"status":  "error",
		})
	}

	if userExists {
		return ctx.JSON(http.StatusConflict, map[string]string{
			"message": "user already exists",
			"status":  "fail",
		})
	}

	isVitian := strings.HasSuffix(payload.Email, "vitstudent.ac.in")

	hashed, err := bcrypt.GenerateFromPassword([]byte(payload.Password), 10)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
			"status":  "error",
		})
	}

	user := models.User{
		ID:         uuid.New(),
		FirstName:  payload.FirstName,
		LastName:   payload.LastName,
		Email:      payload.Email,
		Password:   string(hashed),
		Phone:      payload.Phone,
		College:    payload.College,
		Gender:     payload.Gender,
		Role:       "user",
		Country:    payload.Country,
		Github:     payload.Github,
		Bio:        payload.Bio,
		IsBanned:   false,
		IsAdded:    false,
		IsVitian:   isVitian,
		IsVerified: false,
		TeamID:     0,
	}

	otp, err := utils.GenerateOTP(6)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"status":  "error",
			"message": err.Error(),
		})
	}

	if err := services.InsertUser(user); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
			"status":  "error",
		})
	}

	if err := database.RedisClient.Set(fmt.Sprintf("verfication:%s", user.Email), otp, time.Minute*5); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"status":  "error",
			"message": err.Error(),
		})
	}

	go func() {
		if err := utils.SendMail(user.Email, otp); err != nil {
			slog.Error("error sending email: " + err.Error())
		}
	}()

	return ctx.JSON(http.StatusOK, map[string]string{
		"message": "user creation was successful",
		"status":  "success",
	})
}

func VerifyUser(ctx echo.Context) error {
	var payload models.VerifyUserRequest

	if err := ctx.Bind(&payload); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
			"status":  "fail",
		})
	}

	if err := ctx.Validate(&payload); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
			"status":  "fail",
		})
	}

	user, err := services.FindUserByEmail(payload.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ctx.JSON(http.StatusNotFound, map[string]string{
				"message": "User does not exist",
				"status":  "fail",
			})
		}
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
			"status":  "error",
		})
	}

	if user.IsVerified {
		return ctx.JSON(http.StatusAlreadyReported, map[string]string{
			"message": "user already verified",
			"status":  "success",
		})
	}

	otp, err := database.RedisClient.Get("verification:" + user.Email)
	if err != nil {
		if err == redis.Nil {
			return ctx.JSON(http.StatusForbidden, map[string]string{
				"message": "otp expired",
				"status":  "fail",
			})
		}
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
			"status":  "error",
		})
	}

	if otp != payload.OTP {
		return ctx.JSON(http.StatusUnauthorized, map[string]string{
			"message": "Invalid OTP",
			"status":  "fail",
		})
	}

	user.IsVerified = true

	if err := services.UpdateUser(user); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
			"status":  "error",
		})
	}

	return ctx.JSON(http.StatusOK, map[string]string{
		"message": "User verified",
		"status":  "success",
	})
}
