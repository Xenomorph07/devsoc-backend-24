package controllers

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/CodeChefVIT/devsoc-backend-24/internal/models"
	services "github.com/CodeChefVIT/devsoc-backend-24/internal/services/team"
	"github.com/CodeChefVIT/devsoc-backend-24/internal/utils"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
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

	user := ctx.Get("user").(*models.User)
	if user.TeamID != uuid.Nil {
		return ctx.JSON(http.StatusExpectationFailed, map[string]string{
			"message": "user is already in a team",
			"status":  "failed to create team",
		})
	}

	code := utils.GenerateUniqueTeamCode()

	team := models.Team{
		ID:       uuid.New(),
		Name:     payload.Name,
		Code:     code,
		Round:    0,
		LeaderID: user.ID,
	}

	if err := services.CreateTeam(team); err != nil {
		var pgerr *pgconn.PgError
		if errors.As(err, &pgerr) {
			if pgerr.Code == "23505" {
				return ctx.JSON(http.StatusConflict, map[string]string{
					"message": "team name already exists",
					"status":  "failed to create team",
				})
			}
		}
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
	var user = ctx.Get("user").(*models.User)

	if user.TeamID == uuid.Nil {
		return ctx.JSON(http.StatusExpectationFailed, map[string]string{
			"message": "The user is not in a team",
			"status":  "false",
		})
	}

	team, err := services.FindTeamByTeamID(user.TeamID)

	if err != nil {
		if err == sql.ErrNoRows {
			return ctx.JSON(http.StatusConflict, map[string]string{
				"message": "The user team id does not exist",
				"status":  "false",
			})
		}
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

	user := ctx.Get("user").(*models.User)

	if user.TeamID != uuid.Nil {
		return ctx.JSON(http.StatusExpectationFailed, map[string]string{
			"message": "user is already in a team",
			"status":  "failed to join team",
		})
	}

	team, err := services.FindTeamByCode(payload.Code)

	if err != nil {
		if err == sql.ErrNoRows {
			return ctx.JSON(http.StatusConflict, map[string]string{
				"message": "team code is invalid",
				"status":  "failed to join team",
			})
		}
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Failed to get team details",
			"status":  "false",
		})
	}

	if services.CheckTeamSize(team.ID) {
		return ctx.JSON(http.StatusFailedDependency, map[string]string{
			"message": "team is full",
			"status":  "failed to join team",
		})
	}

	if err := services.UpdateUserTeamDetails(team.ID, user.Email); err != nil {
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
	var user = ctx.Get("user").(*models.User)

	if err := ctx.Bind(&payload); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
			"status":  "failed to bind body",
		})
	}

	if err := ctx.Validate(&payload); err != nil {
		return err
	}

	team, err := services.FindTeamByTeamID(user.TeamID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
			"status":  "failed to fetch team details",
		})
	}

	if team.LeaderID != user.ID {
		return ctx.JSON(http.StatusForbidden, map[string]string{
			"message": "only the leader can kick a member",
			"status":  "failed to kick member",
		})
	}

	if !services.CheckUserInTeam(payload.UserEmail, user.TeamID) {
		return ctx.JSON(http.StatusExpectationFailed, map[string]string{
			"message": "user is not in the team",
			"status":  "failed to kick member",
		})
	}

	if err := services.UpdateUserTeamDetails(uuid.Nil, payload.UserEmail); err != nil {
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
	var user = ctx.Get("user").(*models.User)

	if user.TeamID == uuid.Nil {
		return ctx.JSON(http.StatusExpectationFailed, map[string]string{
			"message": "user is not in the team",
			"status":  "failed to leave team",
		})
	}

	team, err := services.FindTeamByTeamID(user.TeamID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
			"status":  "failed to fetch team details",
		})
	}

	if err := services.UpdateUserTeamDetails(uuid.Nil, user.Email); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
			"status":  "failed to leave team",
		})
	}

	if team.LeaderID == user.ID {
		if err := services.DeleteTeam(user.TeamID); err != nil {
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

	return ctx.JSON(http.StatusOK, map[string]string{
		"message": "team left successfully",
		"status":  "success",
	})
}
