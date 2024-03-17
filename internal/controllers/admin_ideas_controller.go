package controllers

import (
	"net/http"

	services "github.com/CodeChefVIT/devsoc-backend-24/internal/services/idea"
	"github.com/google/uuid"
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

func ShortList(ctx echo.Context) error {
	var req struct {
		TeamID uuid.UUID `json:"team_id"`
	}

	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": "send better json",
			"status":  "fail",
		})
	}

	if req.TeamID == uuid.Nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": "send better json",
			"status":  "fail",
		})
	}

	err := services.ShortlistIdea(req.TeamID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": "db go brrrr : " + err.Error(),
			"status":  "error",
		})
	}

	return ctx.JSON(http.StatusAccepted, map[string]string{
		"message": "the team id has been selected",
		"status":  "success",
	})
}
