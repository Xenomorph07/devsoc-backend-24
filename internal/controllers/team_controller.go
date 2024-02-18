package controllers

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	"github.com/CodeChefVIT/devsoc-backend-24/internal/models"
	services "github.com/CodeChefVIT/devsoc-backend-24/internal/services/team"
	"github.com/CodeChefVIT/devsoc-backend-24/internal/utils"
)

func CreateTeam(ctx echo.Context) error {
	var payload models.CreateTeamRequest

	if err := ctx.Bind(&payload); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
			"status":  "fail",
		})
	}

	if err := ctx.Validate(&payload); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
			"status":  "fail",
		})
	}

	_, err := services.FindTeamByUserID(ctx.Get("user").(models.User).ID)
	if err == nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": "user is already in a team",
			"status":  "fail",
		})
	}

	_, err = services.FindTeamByName(payload.Name)
	if err == nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": "team name already exists",
			"status":  "fail",
		})
	}

	code, err := utils.GenerateUniqueTeamCode()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
			"status":  "error",
		})
	}

	team := models.Team{
		ID:       uuid.New(),
		Name:     payload.Name,
		Code:     code,
		Round:    0,
		LeaderID: ctx.Get("user").(models.User).ID,
		Users:    []models.User{ctx.Get("user").(models.User)},
		Idea:     models.Idea{},
		Project:  models.Project{},
	}

	if err := services.CreateTeam(team); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
			"status":  "error",
		})
	}

	return ctx.JSON(http.StatusCreated, map[string]string{
		"message": "team created successfully",
		"status":  "success",
	})
}

func GetTeamDetails(ctx echo.Context) error {
	userID := ctx.Get("user").(models.User).ID

	team, err := services.FindTeamByUserID(userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ctx.JSON(http.StatusNotFound, map[string]string{
				"status":  "fail",
				"message": "user is not in a team",
			})
		}
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
			"status":  "error",
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "team details fetched successfully",
		"status":  "success",
		"team":    team,
	})
}

func JoinTeam(ctx echo.Context) error {
	var payload models.JoinTeamRequest

	if err := ctx.Bind(&payload); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
			"status":  "fail",
		})
	}

	if err := ctx.Validate(&payload); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
			"status":  "fail",
		})
	}

	_, err := services.FindTeamByUserID(ctx.Get("user").(models.User).ID)
	if err == nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": "user is already in a team",
			"status":  "fail",
		})
	}

	team, err := services.FindTeamByCode(payload.Code)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": "team code is invalid",
			"status":  "fail",
		})
	}

	if len(team.Users) >= 4 {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": "team is full",
			"status":  "fail",
		})
	}

	team.Users = append(team.Users, ctx.Get("user").(models.User))

	if err := services.UpdateTeam(*team); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
			"status":  "error",
		})
	}

	return ctx.JSON(http.StatusOK, map[string]string{
		"message": "team joined successfully",
		"status":  "success",
	})
}

func KickMember(ctx echo.Context) error {
	var payload models.KickMemberRequest
	userID := ctx.Get("user").(models.User).ID

	if err := ctx.Bind(&payload); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
			"status":  "fail",
		})
	}

	if err := ctx.Validate(&payload); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
			"status":  "fail",
		})
	}

	team, err := services.FindTeamByUserID(userID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
			"status":  "error",
		})
	}

	if team.LeaderID != userID {
		return ctx.JSON(http.StatusForbidden, map[string]string{
			"message": "only the leader can kick a member",
			"status":  "fail",
		})
	}

	var isValid bool = false
	for i, user := range team.Users {
		if user.ID == payload.UserID {
			team.Users = append(team.Users[:i], team.Users[i+1:]...)
			isValid = true
			break
		}
	}

	if !isValid {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": "user is not in the team",
			"status":  "fail",
		})
	}

	if err := services.UpdateTeam(*team); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
			"status":  "error",
		})
	}

	return ctx.JSON(http.StatusOK, map[string]string{
		"message": "member kicked successfully",
		"status":  "success",
	})
}

func LeaveTeam(ctx echo.Context) error {
	userID := ctx.Get("user").(models.User).ID

	team, err := services.FindTeamByUserID(userID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
			"status":  "fail",
		})
	}

	if team.LeaderID == userID {
		if err := services.DeleteTeam(team.ID); err != nil {
			return ctx.JSON(http.StatusInternalServerError, map[string]string{
				"message": err.Error(),
				"status":  "error",
			})
		}

		return ctx.JSON(http.StatusOK, map[string]string{
			"message": "team deleted successfully",
			"status":  "success",
		})
	}

	var isValid bool = false
	for i, user := range team.Users {
		if user.ID == userID {
			team.Users = append(team.Users[:i], team.Users[i+1:]...)
			isValid = true
			break
		}
	}

	if !isValid {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": "user is not in the team",
			"status":  "fail",
		})
	}

	if err := services.UpdateTeam(*team); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
			"status":  "error",
		})
	}

	return ctx.JSON(http.StatusOK, map[string]string{
		"message": "team left successfully",
		"status":  "success",
	})
}
