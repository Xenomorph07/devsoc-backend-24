package routes

import (
	"github.com/CodeChefVIT/devsoc-backend-24/internal/controllers"
	"github.com/CodeChefVIT/devsoc-backend-24/internal/middleware"
	"github.com/labstack/echo/v4"
)

func UserRoutes(incomingRoutes *echo.Echo) {
	incomingRoutes.POST("/signup", controllers.CreateUser)

	user := incomingRoutes.Group("/user")
	user.Use(middleware.Protected())
}
