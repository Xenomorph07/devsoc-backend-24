package controllers

import (
	"net/http"

	services "github.com/CodeChefVIT/devsoc-backend-24/internal/services/idea"
	"github.com/labstack/echo/v4"
)

func GetAllIdeas(ctx echo.Context) error {
	project, err := services.GetAllIdeas()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Failed to fetch projects",
			"data":    err.Error(),
			"status":  "false",
		})
	}

	return ctx.JSON(http.StatusAccepted, map[string]interface{}{
		"message": "Successfully fetched projects",
		"data":    project,
		"status":  "true",
	})
}
