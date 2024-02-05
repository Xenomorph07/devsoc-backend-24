package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateUser(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, map[string]string{
		"message": "user creation was successful",
		"status":  "success",
	})
}
