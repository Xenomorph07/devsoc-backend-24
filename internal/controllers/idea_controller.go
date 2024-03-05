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

type response struct {
	Message string      `json:"message"`
	Status  bool        `json:"status"`
	Data    interface{} `json:"data"`
}

func GetIdea(ctx echo.Context) error {
	user := ctx.Get("user").(*models.User)
	teamid := user.TeamID

	if user.TeamID == uuid.Nil {
		return ctx.JSON(http.StatusConflict, response{
			Message: "The user is not in a team",
			Status:  false,
			Data:    &models.Idea{},
		})
	}

	idea, err := services.GetIdeaByTeamID(teamid)
	if err != nil {
		if err == sql.ErrNoRows {
			return ctx.JSON(http.StatusExpectationFailed, map[string]string{
				"message": "Failed to get idea could be cause the user has not made an idea",
				"status":  "fail",
			})
		}
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
			"status":  "error",
		})

	}

	return ctx.JSON(http.StatusAccepted, response{
		Message: "Successfully got the user details",
		Data:    idea,
		Status:  true,
	})
}

func CreateIdea(ctx echo.Context) error {
	var req models.IdeaRequest

	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, response{
			Message: "Failed to parse the data",
			Status:  false,
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
				return ctx.JSON(http.StatusExpectationFailed, response{
					Message: "The team already has an idea",
					Status:  false,
				})
			}
			return ctx.JSON(http.StatusInternalServerError, response{
				Message: "Failed to create the DB entry : " + err.Error(),
				Status:  false,
			})
		}
	}

	return ctx.JSON(http.StatusAccepted, response{
		Message: "Idea has been created",
		Status:  true,
	})
}

func UpdateIdea(ctx echo.Context) error {
	var req models.IdeaRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, response{
			Message: "Failed to parse the data",
			Status:  false,
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
		return ctx.JSON(http.StatusConflict, response{
			Message: "The user is not in a team",
			Status:  false,
		})
	}

	err := services.UpdateIdea(req, user.TeamID)

	if err != nil {
		if errors.Is(err, errors.New("invalid teamid")) {
			return ctx.JSON(http.StatusExpectationFailed, response{
				Message: "The team has not created an idea",
				Status:  false,
			})
		}
		return ctx.JSON(http.StatusInternalServerError, response{
			Message: "Failed to update the idea " + err.Error(),
			Status:  false,
		})
	}

	return ctx.JSON(http.StatusAccepted, response{
		Message: "Idea has been successfully updated",
		Status:  true,
	})
}
