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
		return ctx.JSON(http.StatusInternalServerError, response{
			Message: "Failed to fetch projects",
			Data:    err.Error(),
			Status:  false,
		})
	}

	return ctx.JSON(http.StatusAccepted, response{
		Message: "Successfully fetched projects",
		Data:    project,
		Status:  false,
	})
}

func GetProjectByID(ctx echo.Context) error {
	projectIDParam := ctx.Param("id")
	projectID, err := uuid.Parse(projectIDParam)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response{
			Message: "Invalid ID format",
			Data:    err.Error(),
			Status:  false,
		})
	}
	project, err := services.GetProjectByID(projectID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response{
			Message: "Failed to fetch project",
			Data:    err.Error(),
			Status:  false,
		})
	}
	return ctx.JSON(http.StatusAccepted, response{
		Message: "Successfully fetched project",
		Data:    project,
		Status:  false,
	})
}
