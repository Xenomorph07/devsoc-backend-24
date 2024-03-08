package controllers

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/labstack/echo/v4"

	"github.com/CodeChefVIT/devsoc-backend-24/internal/models"
	services "github.com/CodeChefVIT/devsoc-backend-24/internal/services/projects"
	"github.com/CodeChefVIT/devsoc-backend-24/internal/utils"
)

func GetProject(ctx echo.Context) error {
	user := ctx.Get("user").(*models.User)

	if user.TeamID == uuid.Nil {
		return ctx.JSON(http.StatusConflict, map[string]string{
			"message": "The user is not in a team",
			"status":  "fail",
		})
	}

	proj, err := services.GetProject(user.TeamID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ctx.JSON(http.StatusNotFound, map[string]string{
				"message": "the team does not have a project",
				"status":  "fail",
			})
		}
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Failed to get project : " + err.Error(),
			"status":  "fail",
		})
	}

	return ctx.JSON(http.StatusAccepted, map[string]interface{}{
		"message": "Successfully retrived the project",
		"status":  "success",
		"data":    proj,
	})
}

func CreateProject(ctx echo.Context) error {
	var req models.ProjectRequest

	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": "Failed to parse the data",
			"status":  "fail",
		})
	}

	if err := ctx.Validate(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
			"status":  "fail",
		})
	}

	user := ctx.Get("user").(*models.User)

	if !user.IsLeader {
		return ctx.JSON(http.StatusUnauthorized, map[string]string{
			"message": "user is not a leader",
			"status":  "fail",
		})
	}

	if user.TeamID == uuid.Nil {
		return ctx.JSON(http.StatusForbidden, map[string]string{
			"message": "The user is not in a team",
			"status":  "fail",
		})
	}

	err := services.CreateProject(req, user.TeamID)
	if err != nil {
		var pgerr *pgconn.PgError
		if errors.As(err, &pgerr) {
			if pgerr.Code == "23505" {
				return ctx.JSON(http.StatusConflict, map[string]string{
					"message": "project alread present",
					"status":  "fail",
				})
			}
		}
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Failed to create the project " + err.Error(),
			"status":  "error",
		})
	}

	return ctx.JSON(http.StatusAccepted, map[string]string{
		"message": "Project successfully created",
		"status":  "success",
	})
}

func UpdateProject(ctx echo.Context) error {
	var req models.ProjectRequest

	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": "Failed to parse the data",
			"status":  "fail",
		})
	}

	if err := ctx.Validate(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
			"status":  "fail",
		})
	}

	user := ctx.Get("user").(*models.User)

	if !user.IsLeader {
		return ctx.JSON(http.StatusUnauthorized, map[string]string{
			"message": "user is not a leader",
			"status":  "fail",
		})
	}

	if user.TeamID == uuid.Nil {
		return ctx.JSON(http.StatusForbidden, map[string]string{
			"message": "The user is not in a team",
			"status":  "fail",
		})
	}

	err := services.UpdateProject(req, user.TeamID)
	if err != nil {
		if errors.Is(err, utils.ErrInvalidTeamID) {
			return ctx.JSON(http.StatusNotFound, map[string]string{
				"message": "project not found",
				"status":  "fail",
			})
		}
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Failed to update the project" + err.Error(),
			"status":  "error",
		})
	}

	return ctx.JSON(http.StatusAccepted, map[string]string{
		"message": "Project successfully updated",
		"status":  "success",
	})
}
