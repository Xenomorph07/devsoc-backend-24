package routes

import (
	"github.com/CodeChefVIT/devsoc-backend-24/internal/controllers"
	"github.com/labstack/echo/v4"
)

func UserRoutes(incomingRoutes *echo.Echo) {
	user := incomingRoutes.Group("/user")
	user.POST("/create", controllers.CreateUser)
}
