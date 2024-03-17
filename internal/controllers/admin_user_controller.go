package controllers

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/CodeChefVIT/devsoc-backend-24/internal/models"
	admin "github.com/CodeChefVIT/devsoc-backend-24/internal/services/admin"
	services "github.com/CodeChefVIT/devsoc-backend-24/internal/services/user"
)

func GetAllUsers(ctx echo.Context) error {
	users, err := admin.GetAllUsers()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Failed to fetch users",
			"error":   err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "Users fetched successfully",
		"users":   users,
	})
}

func GetAllVitians(ctx echo.Context) error {
	users, err := admin.GetAllVitians()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Failed to fetch users",
			"error":   err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "Users fetched successfully",
		"users":   users,
	})
}

func GetAllFemales(ctx echo.Context) error {
	users, err := admin.GetAllFemales()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Failed to fetch users",
			"error":   err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "Users fetched successfully",
		"users":   users,
	})
}

func GetUserByEmail(ctx echo.Context) error {
	email := ctx.Param("email")
	user, err := services.FindUserByEmail(email)
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
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "User fetched successfully",
		"user":    user,
	})

}

func BanUser(ctx echo.Context) error {
	var payload models.BanUser

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

	user.IsBanned = true

	err = services.UpdateUser(&user.User)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
			"status":  "fail",
		})
	}

	return ctx.JSON(http.StatusOK, map[string]string{
		"message": "User banned",
		"status":  "success",
	})
}

func UnbanUser(ctx echo.Context) error {
	var payload models.BanUser

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

	user.IsBanned = false

	err = services.UpdateUser(&user.User)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
			"status":  "fail",
		})
	}

	return ctx.JSON(http.StatusOK, map[string]string{
		"message": "User unbanned",
		"status":  "success",
	})
}

func CheckIn(ctx echo.Context) error {
	var payload struct {
		Email string `json:"email" validate:"required"`
	}

	if err := ctx.Bind(&payload); err != nil {
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

	user.IsAdded = true

	err = services.UpdateUser(&user.User)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
			"status":  "fail",
		})
	}

	return ctx.JSON(http.StatusOK, map[string]string{
		"message": "checkin successful",
		"status":  "success",
	})
}
