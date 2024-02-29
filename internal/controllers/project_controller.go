package controllers

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/CodeChefVIT/devsoc-backend-24/internal/models"
	services "github.com/CodeChefVIT/devsoc-backend-24/internal/services/projects"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/labstack/echo/v4"
)

func GetProject(ctx echo.Context) error {

	user := ctx.Get("user").(*models.User)

	if user.TeamID == uuid.Nil {
		return ctx.JSON(http.StatusConflict, response{
			Message: "The user is not in a team",
			Status:  false,
			Data:    models.GetProject{},
		})
	}

	proj, err := services.GetProject(user.TeamID)
	if err != nil {
		if err == sql.ErrNoRows {
			return ctx.JSON(http.StatusExpectationFailed, response{
				Message: "Failed to get project could be cause the user has not made an idea",
				Data:    models.GetProject{},
				Status:  false,
			})
		}
		return ctx.JSON(http.StatusInternalServerError, response{
			Message: "Failed to get project : " + err.Error(),
			Status:  false,
			Data:    models.GetProject{},
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

	if user.TeamID == uuid.Nil {
		return ctx.JSON(http.StatusConflict, response{
			Message: "The user is not in a team",
			Status:  false,
		})
	}

	err := services.CreateProject(req, user.TeamID)
	if err != nil {
		var pgerr *pgconn.PgError
		if errors.As(err, &pgerr) {
			if pgerr.Code == "23505" {
				return ctx.JSON(http.StatusExpectationFailed, response{
					Message: "The team already has an project",
					Status:  false,
				})
			}
		}
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

	if user.TeamID == uuid.Nil {
		return ctx.JSON(http.StatusConflict, response{
			Message: "The user is not in a team",
			Status:  false,
		})
	}

	err := services.UpdateProject(req, user.TeamID)
	if err != nil {
		if errors.Is(err, errors.New("invalid teamid")) {
			return ctx.JSON(http.StatusExpectationFailed, response{
				Message: "The team has not created an project",
				Status:  false,
			})
		}
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
