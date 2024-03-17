package routes

import (
	"github.com/CodeChefVIT/devsoc-backend-24/internal/controllers"
	"github.com/CodeChefVIT/devsoc-backend-24/internal/middleware"
	"github.com/labstack/echo/v4"
)

func AdminRoutes(incomingRoutes *echo.Echo) {
	admin := incomingRoutes.Group("/admin")
	admin.Use(middleware.Protected())
	admin.Use(middleware.CheckAdmin)

	admin.GET("/users", controllers.GetAllUsers)
	admin.GET("/user/:email", controllers.GetUserByEmail)
	admin.POST("/user/ban", controllers.BanUser, middleware.EditOnly)
	admin.POST("/user/unban", controllers.UnbanUser, middleware.EditOnly)
	admin.GET("/vitians", controllers.GetAllVitians)
	admin.GET("/females", controllers.GetAllFemales)

	admin.GET("/team/all", controllers.GetTeams)
	admin.GET("/team/:id", controllers.GetTeamsByID)
	admin.GET("/team/project/:id", controllers.GetProjectByTeamID)
	admin.GET("/team/leader/:id", controllers.GetTeamLeader)
	admin.GET("/team/idea/:id", controllers.GetIdeaByTeamID)
	admin.POST("/team/ban/:id", controllers.BanTeam, middleware.EditOnly)
	admin.POST("/team/unban/:id", controllers.UnbanTeam, middleware.EditOnly)

	admin.GET("/projects/all", controllers.GetAllProject)
	admin.GET("/ideas/all", controllers.GetAllIdeas)

	admin.POST("/ideas/shortlist", controllers.ShortList)

	//admin.GET("/team/freshers", controllers.GetAllFresherTeams)
	//admin.GET("/team/females", controllers.GetAllFemaleTeams)
}
