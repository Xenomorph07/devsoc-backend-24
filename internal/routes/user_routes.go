package routes

import (
	"github.com/labstack/echo/v4"

	"github.com/CodeChefVIT/devsoc-backend-24/internal/controllers"
)

func UserRoutes(incomingRoutes *echo.Echo) {
	incomingRoutes.POST("/signup", controllers.CreateUser)
	incomingRoutes.POST("/verify", controllers.VerifyUser)
	incomingRoutes.POST("/resend", controllers.ResendOTP)

	user := incomingRoutes.Group("/user")
	user.POST("/complete-profile", controllers.CompleteProfile)

}
