package routes

import (
	"github.com/CodeChefVIT/devsoc-backend-24/internal/controllers"
	"github.com/CodeChefVIT/devsoc-backend-24/internal/middleware"
	"github.com/labstack/echo/v4"
)

func AuthRoutes(incomingRoutes *echo.Echo) {
	incomingRoutes.POST("/login", controllers.Login)
	incomingRoutes.POST("/logout", controllers.Logout, middleware.Refresh())
	incomingRoutes.POST("/refresh", controllers.Refresh, middleware.Refresh())
}
