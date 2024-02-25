package controllers

import (
	"net/http"

	"github.com/CodeChefVIT/devsoc-backend-24/internal/models"
	services "github.com/CodeChefVIT/devsoc-backend-24/internal/services/idea"
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

	idea, err := services.GetIdeaByTeamID(teamid)
	if err != nil {
		return ctx.JSON(http.StatusExpectationFailed, response{
			Message: "Failed to get idea could be cause the user has not made an idea",
			Data:    idea,
			Status:  false,
		})
	}

	return ctx.JSON(http.StatusAccepted, response{
		Message: "Successfully got the user details",
		Data:    idea,
		Status:  true,
	})
}

func CreateIdea(ctx echo.Context) error {
	var req models.CreateUpdateIdeasRequest

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

	err := services.CreateIdea(req, user.TeamID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response{
			Message: "Failed to create the DB entry : " + err.Error(),
			Status:  false,
		})
	}

	return ctx.JSON(http.StatusAccepted, response{
		Message: "Idea has been created",
		Status:  true,
	})
}

func UpdateIdea(ctx echo.Context) error {

	var req models.CreateUpdateIdeasRequest
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

	err := services.UpdateIdea(req, user.TeamID)

	if err != nil {
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
