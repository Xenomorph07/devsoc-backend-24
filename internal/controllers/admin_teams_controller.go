package controllers

import (
	"net/http"

	ideaService "github.com/CodeChefVIT/devsoc-backend-24/internal/services/idea"
	projectService "github.com/CodeChefVIT/devsoc-backend-24/internal/services/projects"
	teamService "github.com/CodeChefVIT/devsoc-backend-24/internal/services/team"
	userService "github.com/CodeChefVIT/devsoc-backend-24/internal/services/user"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func GetTeamsByID(ctx echo.Context) error {
	teamIDParam := ctx.Param("id")
	teamID, err := uuid.Parse(teamIDParam)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": "Invalid team ID format",
			"status":  "false",
			"data":    err.Error(),
		})
	}

	team, err := teamService.FindTeamByTeamID(teamID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Failed to fetch team",
			"data":    err.Error(),
			"status":  "false",
		})
	}

	return ctx.JSON(http.StatusAccepted, map[string]interface{}{
		"message": "Successfully got Team",
		"data":    team,
		"status":  "true",
	})
}

func GetIdeaByTeamID(ctx echo.Context) error {
	teamIDParam := ctx.Param("id")
	teamID, err := uuid.Parse(teamIDParam)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": "Invalid team ID format",
			"status":  "false",
			"data":    err.Error(),
		})
	}

	team, err := ideaService.GetIdeaByTeamID(teamID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Failed to fetch ideas",
			"data":    err.Error(),
			"status":  "false",
		})
	}

	return ctx.JSON(http.StatusAccepted, map[string]interface{}{
		"message": "Successfully got Idea",
		"data":    team,
		"status":  "true",
	})
}

func GetProjectByTeamID(ctx echo.Context) error {
	teamIDParam := ctx.Param("id")
	teamID, err := uuid.Parse(teamIDParam)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": "Invalid team ID format",
			"status":  "false",
			"data":    err.Error(),
		})
	}

	team, err := projectService.GetProject(teamID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Failed to get project",
			"data":    err.Error(),
			"status":  "false",
		})
	}

	return ctx.JSON(http.StatusAccepted, map[string]interface{}{
		"message": "Successfully got Project",
		"data":    team,
		"status":  "true",
	})
}

func GetTeams(ctx echo.Context) error {
	team, err := teamService.GetAllTeams()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Failed to fetch teams",
			"data":    err.Error(),
			"status":  "false",
		})
	}
	return ctx.JSON(http.StatusAccepted, map[string]interface{}{
		"message": "Successfully fetched teams",
		"data":    team,
		"status":  "false",
	})
}

func BanTeam(ctx echo.Context) error {
	teamIDParam := ctx.Param("id")
	teamID, err := uuid.Parse(teamIDParam)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": "Failed to ban user",
			"data":    err.Error(),
			"status":  "false",
		})
	}
	err = teamService.BanTeam(teamID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Failed to ban team",
			"data":    err.Error(),
			"status":  "false",
		})
	}
	return ctx.JSON(http.StatusAccepted, map[string]string{
		"message": "Successfully banned Team",
		"status":  "true",
	})
}
func UnbanTeam(ctx echo.Context) error {
	teamIDParam := ctx.Param("id")
	teamID, err := uuid.Parse(teamIDParam)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": "Invalid ID format",
			"data":    err.Error(),
			"status":  "false",
		})
	}
	err = teamService.UnbanTeam(teamID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Failed to unban team",
			"data":    err.Error(),
			"status":  "false",
		})
	}
	return ctx.JSON(http.StatusAccepted, map[string]string{
		"message": "Successfully unbanned Team",
		"status":  "true",
	})
}

func GetTeamLeader(ctx echo.Context) error {
	teamIDParam := ctx.Param("id")
	teamID, err := uuid.Parse(teamIDParam)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": "Invalid team ID format",
			"status":  "false",
			"data":    err.Error(),
		})
	}
	team, err := teamService.FindTeamByTeamID(teamID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Failed to fetch team",
			"data":    err.Error(),
			"status":  "false",
		})
	}
	user, err := userService.FindUserByID(team.LeaderID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Failed to fetch user",
			"data":    err.Error(),
			"status":  "false",
		})
	}
	return ctx.JSON(http.StatusAccepted, map[string]interface{}{
		"message": "Successfully fetched leader",
		"data":    user,
		"status":  "true",
	})

}

/*func GetAllFresherTeams(ctx echo.Context) error {
	team, err := teamService.GetAllFresherTeam()
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": "Failed to fetch teams",
			"data":    err.Error(),
			"status":  "false",
		})
	}
	return ctx.JSON(http.StatusAccepted, map[string]interface{}{
		"message": "Successfully fetched teams",
		"data":    team,
		"status":  "true",
	})
}

func GetAllFemaleTeams(ctx echo.Context) error {
	team, err := teamService.GetAllFemaleTeams()
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": "Failed to fetch teams",
			"data":    err.Error(),
			"status":  "false",
		})
	}
	return ctx.JSON(http.StatusAccepted, map[string]interface{}{
		"message": "Successfully fetched teams",
		"data":    team,
		"status":  "true",
	})
}*/
