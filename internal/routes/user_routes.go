package routes

import (
	"github.com/labstack/echo/v4"

	"github.com/CodeChefVIT/devsoc-backend-24/internal/controllers"
	"github.com/CodeChefVIT/devsoc-backend-24/internal/middleware"
)

func UserRoutes(incomingRoutes *echo.Echo) {
	incomingRoutes.POST("/signup", controllers.CreateUser)
	incomingRoutes.POST("/verify", controllers.VerifyUser)
	incomingRoutes.POST("/resend", controllers.ResendOTP)
	incomingRoutes.POST("/reset-password", controllers.RequestResetPassword)
	incomingRoutes.PATCH("/reset-password", controllers.ResetPassword)

	user := incomingRoutes.Group("/user")
	user.POST("/complete-profile", controllers.CompleteProfile)
	user.GET("/me", controllers.Dashboard, middleware.Protected(), middleware.AuthUser)
}
