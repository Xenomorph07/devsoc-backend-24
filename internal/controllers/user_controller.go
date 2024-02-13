package controllers

import (
	"database/sql"
	"errors"
	"net/http"
	"strings"

	"github.com/CodeChefVIT/devsoc-backend-24/internal/models"
	services "github.com/CodeChefVIT/devsoc-backend-24/internal/services/user"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
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
		return err
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

	if err := services.InsertUser(user); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
			"status":  "error",
		})
	}

	return ctx.JSON(http.StatusOK, map[string]string{
		"message": "user creation was successful",
		"status":  "success",
	})
}
