package controllers

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/CodeChefVIT/devsoc-backend-24/internal/models"
	services "github.com/CodeChefVIT/devsoc-backend-24/internal/services/idea"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/labstack/echo/v4"
)

func GetIdea(ctx echo.Context) error {
	user := ctx.Get("user").(*models.User)
	teamid := user.TeamID

	if user.TeamID == uuid.Nil {
		return ctx.JSON(http.StatusConflict, map[string]string{
			"message": "The user is not in a team",
			"status":  "fail",
		})
	}

	idea, err := services.GetIdeaByTeamID(teamid)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ctx.JSON(http.StatusNotFound, map[string]string{
				"message": "the team has not made an idea yet",
				"status":  "fail",
			})
		}
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
			"status":  "error",
		})

	}

	return ctx.JSON(http.StatusAccepted, map[string]interface{}{
		"message": "idea found",
		"data":    idea,
		"status":  "success",
	})
}

func CreateIdea(ctx echo.Context) error {
	var req models.IdeaRequest

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

	err := services.CreateIdea(req, user.TeamID)
	if err != nil {
		var pgerr *pgconn.PgError
		if errors.As(err, &pgerr) {
			if pgerr.Code == "23505" {
				return ctx.JSON(http.StatusConflict, map[string]string{
					"message": "The team already has an idea",
					"status":  "fail",
				})
			}
			return ctx.JSON(http.StatusInternalServerError, map[string]string{
				"message": "Failed to create the DB entry : " + err.Error(),
				"status":  "error",
			})
		}
	}

	return ctx.JSON(http.StatusAccepted, map[string]string{
		"message": "idea has been created",
		"status":  "success",
	})
}

func UpdateIdea(ctx echo.Context) error {
	var req models.IdeaRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": "Failed to parse the data",
			"status":  "error",
		})
	}

	if err := ctx.Validate(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
			"status":  "fail",
		})
	}

	user := ctx.Get("user").(*models.User)

	if user.TeamID == uuid.Nil {
		return ctx.JSON(http.StatusConflict, map[string]string{
			"message": "The user is not in a team",
			"status":  "fail",
		})
	}

	err := services.UpdateIdea(req, user.TeamID)

	if err != nil {
		if errors.Is(err, errors.New("invalid teamid")) {
			return ctx.JSON(http.StatusNotFound, map[string]string{
				"message": "idea not found",
				"status":  "fail",
			})
		}
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Failed to update the idea " + err.Error(),
			"status":  "error",
		})
	}

	return ctx.JSON(http.StatusAccepted, map[string]string{
		"message": "idea successfully updated",
		"status":  "success",
	})
}
