package routes

import (
	"github.com/CodeChefVIT/devsoc-backend-24/internal/controllers"
	"github.com/CodeChefVIT/devsoc-backend-24/internal/middleware"
	"github.com/labstack/echo/v4"
)

func IdeaRoutes(incomingRoutes *echo.Echo) {
	idea := incomingRoutes.Group("/idea")
	idea.Use(middleware.Protected())
	idea.Use(middleware.AuthUser)
	idea.GET("/", controllers.GetIdea)
	idea.POST("/create", controllers.CreateIdea)
	idea.PATCH("/update", controllers.UpdateIdea)
}
