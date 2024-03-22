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
	incomingRoutes.POST("/admin/login", controllers.AdminLogin)

	user := incomingRoutes.Group("/user")
	user.POST(
		"/complete-profile",
		controllers.CompleteProfile,
		middleware.Protected(),
		middleware.AuthUser,
	)
	user.GET("/me", controllers.Dashboard, middleware.Protected(), middleware.AuthUser)
	user.PATCH("/update", controllers.UpdateUser, middleware.Protected(), middleware.AuthUser)
}
