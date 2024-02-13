package routes

import (
	"github.com/CodeChefVIT/devsoc-backend-24/internal/controllers"
	"github.com/CodeChefVIT/devsoc-backend-24/internal/middleware"
	"github.com/labstack/echo/v4"
)

func TeamRoutes(incomingRoutes *echo.Echo) {
	team := incomingRoutes.Group("/team")
	team.Use(middleware.Protected(), middleware.AuthUser)
	team.GET("/", controllers.GetTeamDetails)
	team.POST("/create", controllers.CreateTeam)
	team.POST("/join", controllers.JoinTeam)
	team.DELETE("/leave", controllers.LeaveTeam)
	team.POST("/kick", controllers.KickMember)
}
