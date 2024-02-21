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

	userid := ctx.Get("user").(*models.User).ID
	err := services.CheckUserTeam(userid)
	if err == nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": "user is already in a team",
			"status":  "failed to create team",
		})
	}

	if services.CheckTeamName(payload.Name) {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": "team name already exists",
			"status":  "failed to create team",
		})
	}

	code := utils.GenerateUniqueTeamCode()

	team := models.Team{
		ID:       uuid.New(),
		Name:     payload.Name,
		Code:     code,
		Round:    0,
		LeaderID: userid,
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
	var user = ctx.Get("user").(*models.User)

	team, err := services.FindTeamByTeamID(user.TeamID)

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

/*func JoinTeam(ctx echo.Context) error {
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

	user_id := ctx.Get("user").(*models.User).ID

	_, err := services.FindTeamByUserID(user_id)
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

	//team.Users = append(team.Users, user_id)

	if err := services.UpdateUserTeamDetails(user_id); err != nil {
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
	var userID = ctx.Get("user").(*models.User).ID

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

	/*var isValid bool = false
	for i, user := range team.Users {
		if user == payload.UserID {
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

	if err := services.UpdateTeam(team); err != nil {
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
	var userID = ctx.Get("user").(*models.User).ID

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
		if user == userID {
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

	if err := services.UpdateTeam(team); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
			"status":  "failed to leave team",
		})
	}

	return ctx.JSON(http.StatusOK, map[string]string{
		"message": "team left successfully",
		"status":  "success",
	})
}*/
