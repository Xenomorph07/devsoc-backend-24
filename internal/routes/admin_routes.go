package routes

import (
	"github.com/CodeChefVIT/devsoc-backend-24/internal/controllers"
	"github.com/labstack/echo/v4"
)

func AdminRoutes(incomingRoutes *echo.Echo) {
	admin := incomingRoutes.Group("/admin")
	// admin.Use(middleware.Protected())
	// admin.Use(middleware.AuthUser)

	admin.GET("/team/all", controllers.GetTeams)
	admin.GET("/team/:id", controllers.GetTeamsByID)
	admin.GET("/team/project/:id", controllers.GetProjectByTeamID)
	admin.GET("/team/leader/:id", controllers.GetTeamLeader)
	admin.GET("/team/idea/:id", controllers.GetIdeaByTeamID)
	admin.GET("/team/ban/:id", controllers.BanTeam)
	admin.GET("/team/unban/:id", controllers.UnbanTeam)
}
