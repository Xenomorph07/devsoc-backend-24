package routes

import (
	"github.com/CodeChefVIT/devsoc-backend-24/internal/controllers"
	"github.com/CodeChefVIT/devsoc-backend-24/internal/middleware"
	"github.com/labstack/echo/v4"
)

func ProjectRoutes(incomingRoutes *echo.Echo) {
	project := incomingRoutes.Group("/project")
	project.Use(middleware.Protected())
	project.Use(middleware.AuthUser)
	project.GET("", controllers.GetProject)
	project.POST("/create", controllers.CreateProject)
	project.PATCH("/update", controllers.UpdateProject)
}
