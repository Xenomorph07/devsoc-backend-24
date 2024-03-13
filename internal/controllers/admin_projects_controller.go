package controllers

import (
	"net/http"

	services "github.com/CodeChefVIT/devsoc-backend-24/internal/services/projects"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func GetAllProject(ctx echo.Context) error {
	project, err := services.GetAllProjects()
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

func GetProjectByID(ctx echo.Context) error {
	projectIDParam := ctx.Param("id")
	projectID, err := uuid.Parse(projectIDParam)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": "Invalid ID format",
			"data":    err.Error(),
			"status":  "false",
		})
	}
	project, err := services.GetProjectByID(projectID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Failed to fetch project",
			"data":    err.Error(),
			"status":  "false",
		})
	}
	return ctx.JSON(http.StatusAccepted, map[string]interface{}{
		"message": "Successfully fetched project",
		"data":    project,
		"status":  "true",
	})
}
