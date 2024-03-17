package controllers

import (
	"database/sql"
	"errors"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/labstack/echo/v4"

	"github.com/CodeChefVIT/devsoc-backend-24/internal/models"
	services "github.com/CodeChefVIT/devsoc-backend-24/internal/services/projects"
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

	req.Name = strings.TrimSpace(req.Name)
	req.Description = strings.TrimSpace(req.Description)
	req.Track = strings.TrimSpace(req.Track)
	req.GithubLink = strings.TrimSpace(req.GithubLink)
	req.FigmaLink = strings.TrimSpace(req.FigmaLink)
	req.Others = strings.TrimSpace(req.Others)

	if err := ctx.Validate(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
			"status":  "fail",
		})
	}

	user := ctx.Get("user").(*models.User)

	if user.TeamID == uuid.Nil {
		return ctx.JSON(http.StatusConflict, map[string]string{
			"message": "user is not in a team",
			"status":  "fail",
		})
	}

	if !user.IsLeader {
		return ctx.JSON(http.StatusForbidden, map[string]string{
			"message": "user is not a leader",
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
	var req models.UpdateProjectRequest

	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": "Failed to parse the data",
			"status":  "fail",
		})
	}

	req.Name = strings.TrimSpace(req.Name)
	req.Description = strings.TrimSpace(req.Description)
	req.Track = strings.TrimSpace(req.Track)
	req.GithubLink = strings.TrimSpace(req.GithubLink)
	req.FigmaLink = strings.TrimSpace(req.FigmaLink)
	req.Others = strings.TrimSpace(req.Others)

	if err := ctx.Validate(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
			"status":  "fail",
		})
	}

	user := ctx.Get("user").(*models.User)

	if !user.IsLeader {
		return ctx.JSON(http.StatusConflict, map[string]string{
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

	curr, err := services.GetProject(user.TeamID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ctx.JSON(http.StatusNotFound, map[string]string{
				"message": "project not found",
				"status":  "fail",
			})
		}
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": "db error : " + err.Error(),
			"status":  "fail",
		})
	}

	if req.Name == "" {
		req.Name = curr.Name
	}
	if req.Description == "" {
		req.Description = curr.Description
	}
	if req.Track == "" {
		req.Track = curr.Track
	}
	if req.GithubLink == "" {
		req.GithubLink = curr.GithubLink
	}
	if req.FigmaLink == "" {
		req.FigmaLink = curr.FigmaLink
	}
	if req.Others == "" {
		req.Others = curr.Others
	}

	err = services.UpdateProject(req, user.TeamID)
	if err != nil {
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
