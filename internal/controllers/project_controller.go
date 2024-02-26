package controllers

import (
	"net/http"

	"github.com/CodeChefVIT/devsoc-backend-24/internal/models"
	services "github.com/CodeChefVIT/devsoc-backend-24/internal/services/projects"
	"github.com/labstack/echo/v4"
)

func GetProject(ctx echo.Context) error {

	user := ctx.Get("user").(*models.User)
	proj, err := services.GetProject(user.TeamID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response{
			Message: "Failed to get project : " + err.Error(),
			Status:  false,
			Data:    models.Project{},
		})
	}

	return ctx.JSON(http.StatusAccepted, response{
		Message: "Successfully retrived the project",
		Status:  true,
		Data:    proj,
	})
}

func CreateProject(ctx echo.Context) error {
	var req models.CreateUpdateProjectRequest

	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, response{
			Message: "Failed to parse the data",
			Status:  false,
		})
	}

	if err := ctx.Validate(&req); err != nil {
		return err
	}

	user := ctx.Get("user").(*models.User)

	err := services.CreateProject(req, user.TeamID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response{
			Message: "Failed to create the project " + err.Error(),
			Status:  false,
		})
	}

	return ctx.JSON(http.StatusAccepted, response{
		Message: "Project successfully created",
		Status:  true,
	})
}

func UpdateProject(ctx echo.Context) error {
	var req models.CreateUpdateProjectRequest

	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, response{
			Message: "Failed to parse the data",
			Status:  false,
		})
	}

	if err := ctx.Validate(&req); err != nil {
		return err
	}

	user := ctx.Get("user").(*models.User)

	err := services.UpdateProject(req, user.TeamID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response{
			Message: "Failed to update the project" + err.Error(),
			Status:  false,
		})
	}

	return ctx.JSON(http.StatusAccepted, response{
		Message: "Project successfully updated",
		Status:  true,
	})
}
