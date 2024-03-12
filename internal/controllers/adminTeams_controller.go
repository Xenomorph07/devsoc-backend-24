package controllers

import (
	"net/http"

	"github.com/CodeChefVIT/devsoc-backend-24/internal/models"
	servic "github.com/CodeChefVIT/devsoc-backend-24/internal/services/idea"
	service "github.com/CodeChefVIT/devsoc-backend-24/internal/services/projects"
	services "github.com/CodeChefVIT/devsoc-backend-24/internal/services/team"
	servi "github.com/CodeChefVIT/devsoc-backend-24/internal/services/user"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func GetTeamsByID(ctx echo.Context) error {
	teamIDParam := ctx.Param("id")
	teamID, err := uuid.Parse(teamIDParam)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response{
			Message: "Invalid team ID format",
			Status:  false,
			Data:    err.Error(),
		})
	}

	team, err := services.FindTeamByTeamID(teamID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response{
			Message: "Failed to fetch team",
			Data:    err.Error(),
			Status:  false,
		})
	}

	return ctx.JSON(http.StatusAccepted, response{
		Message: "Successfully got Team",
		Data:    team,
		Status:  true,
	})
}

func GetIdeaByTeamID(ctx echo.Context) error {
	teamIDParam := ctx.Param("id")
	teamID, err := uuid.Parse(teamIDParam)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response{
			Message: "Invalid team ID format",
			Status:  false,
			Data:    models.Team{},
		})
	}

	team, err := servic.GetIdeaByTeamID(teamID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response{
			Message: "Failed to fetch ideas",
			Data:    err.Error(),
			Status:  false,
		})
	}

	return ctx.JSON(http.StatusAccepted, response{
		Message: "Successfully got Idea",
		Data:    team,
		Status:  true,
	})
}

func GetProjectByTeamID(ctx echo.Context) error {
	teamIDParam := ctx.Param("id")
	teamID, err := uuid.Parse(teamIDParam)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response{
			Message: "Invalid team ID format",
			Status:  false,
			Data:    models.Team{},
		})
	}

	team, err := service.GetProject(teamID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response{
			Message: err.Error(),
			Data:    team,
			Status:  false,
		})
	}

	return ctx.JSON(http.StatusAccepted, response{
		Message: "Successfully got Project",
		Data:    team,
		Status:  true,
	})
}

func GetTeams(ctx echo.Context) error {
	team, err := services.GetAllTeams()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response{
			Message: "Failed to fetch teams",
			Data:    err.Error(),
			Status:  false,
		})
	}
	return ctx.JSON(http.StatusAccepted, response{
		Message: "Successfully fetched teams",
		Data:    team,
		Status:  false,
	})
}

func BanTeam(ctx echo.Context) error {
	teamIDParam := ctx.Param("id")
	teamID, err := uuid.Parse(teamIDParam)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response{
			Message: err.Error(),
			Data:    models.Team{},
			Status:  false,
		})
	}
	err = services.BanTeam(teamID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response{
			Message: "Failed to ban team",
			Data:    err.Error(),
			Status:  false,
		})
	}
	return ctx.JSON(http.StatusAccepted, response{
		Message: "Successfully banned Team",
		Status:  true,
	})
}
func UnbanTeam(ctx echo.Context) error {
	teamIDParam := ctx.Param("id")
	teamID, err := uuid.Parse(teamIDParam)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response{
			Message: err.Error(),
			Data:    models.Team{},
			Status:  false,
		})
	}
	err = services.UnbanTeam(teamID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response{
			Message: "Failed to unban team",
			Data:    err.Error(),
			Status:  false,
		})
	}
	return ctx.JSON(http.StatusAccepted, response{
		Message: "Successfully unbanned Team",
		Status:  true,
	})
}

func GetTeamLeader(ctx echo.Context) error {
	teamIDParam := ctx.Param("id")
	teamID, err := uuid.Parse(teamIDParam)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response{
			Message: "Invalid team ID format",
			Status:  false,
			Data:    err.Error(),
		})
	}
	team, err := services.FindTeamByTeamID(teamID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response{
			Message: "Failed to fetch team",
			Data:    err.Error(),
			Status:  false,
		})
	}
	user, err := servi.FindUserByID(team.LeaderID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response{
			Message: "Failed to fetch user",
			Data:    err.Error(),
			Status:  false,
		})
	}
	return ctx.JSON(http.StatusAccepted, response{
		Message: "Successfully fetched leader",
		Data:    user,
		Status:  true,
	})

}
