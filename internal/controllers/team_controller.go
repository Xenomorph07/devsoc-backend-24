package controllers

import (
	"net/http"

	"github.com/CodeChefVIT/devsoc-backend-24/internal/models"
	services "github.com/CodeChefVIT/devsoc-backend-24/internal/services/team"
	"github.com/CodeChefVIT/devsoc-backend-24/internal/utils"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func CreateTeam(ctx echo.Context) error {
	var payload models.CreateTeamRequest

	if err := ctx.Bind(&payload); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
			"status":  "failed to bind body",
		})
	}

	if err := ctx.Validate(&payload); err != nil {
		return err
	}

	_, err := services.FindTeamByUserID(ctx.Get("user").(models.User).ID)
	if err == nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": "user is already in a team",
			"status":  "failed to create team",
		})
	}

	_, err = services.FindTeamByName(payload.Name)
	if err == nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": "team name already exists",
			"status":  "failed to create team",
		})
	}

	code, err := utils.GenerateUniqueTeamCode()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
			"status":  "failed to generate team code",
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
			"status":  "failed to create team",
		})
	}

	return ctx.JSON(http.StatusCreated, map[string]string{
		"message": "team created successfully",
		"status":  "success",
	})
}

func GetTeamDetails(ctx echo.Context) error {
	var userID = ctx.Get("user").(models.User).ID

	team, err := services.FindTeamByUserID(userID)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
			"status":  "failed to fetch team details",
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
			"status":  "failed to bind body",
		})
	}

	if err := ctx.Validate(&payload); err != nil {
		return err
	}

	_, err := services.FindTeamByUserID(ctx.Get("user").(models.User).ID)
	if err == nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": "user is already in a team",
			"status":  "failed to join team",
		})
	}

	team, err := services.FindTeamByCode(payload.Code)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": "team code is invalid",
			"status":  "failed to join team",
		})
	}

	if len(team.Users) >= 4 {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": "team is full",
			"status":  "failed to join team",
		})
	}

	team.Users = append(team.Users, ctx.Get("user").(models.User))

	if err := services.UpdateTeam(*team); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
			"status":  "failed to join team",
		})
	}

	return ctx.JSON(http.StatusOK, map[string]string{
		"message": "team joined successfully",
		"status":  "success",
	})
}

func KickMember(ctx echo.Context) error {
	var payload models.KickMemberRequest
	var userID = ctx.Get("user").(models.User).ID

	if err := ctx.Bind(&payload); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
			"status":  "failed to bind body",
		})
	}

	if err := ctx.Validate(&payload); err != nil {
		return err
	}

	team, err := services.FindTeamByUserID(userID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
			"status":  "failed to fetch team details",
		})
	}

	if team.LeaderID != userID {
		return ctx.JSON(http.StatusForbidden, map[string]string{
			"message": "only the leader can kick a member",
			"status":  "failed to kick member",
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
			"status":  "failed to kick member",
		})
	}

	if err := services.UpdateTeam(*team); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
			"status":  "failed to kick member",
		})
	}

	return ctx.JSON(http.StatusOK, map[string]string{
		"message": "member kicked successfully",
		"status":  "success",
	})
}

func LeaveTeam(ctx echo.Context) error {
	var userID = ctx.Get("user").(models.User).ID

	team, err := services.FindTeamByUserID(userID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
			"status":  "failed to fetch team details",
		})
	}

	if team.LeaderID == userID {
		if err := services.DeleteTeam(team.ID); err != nil {
			return ctx.JSON(http.StatusInternalServerError, map[string]string{
				"message": err.Error(),
				"status":  "failed to leave team",
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
			"status":  "failed to leave team",
		})
	}

	if err := services.UpdateTeam(*team); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
			"status":  "failed to leave team",
		})
	}

	return ctx.JSON(http.StatusOK, map[string]string{
		"message": "team left successfully",
		"status":  "success",
	})
}
