package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"

	"github.com/CodeChefVIT/devsoc-backend-24/config"
	"github.com/CodeChefVIT/devsoc-backend-24/internal/database"
	"github.com/CodeChefVIT/devsoc-backend-24/internal/routes"
	"github.com/CodeChefVIT/devsoc-backend-24/internal/utils"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func init() {
	config.SanityCheck()
	appConfig := config.LoadConfig()
	database.InitDB(appConfig.DatabaseConfig)
	database.InitialiseGoogleSheetsClient()
	database.InitRedis(appConfig.RedisConfig)
}

func main() {
	app := echo.New()
	app.Validator = &utils.Validator{Validator: validator.New(validator.WithRequiredStructEnabled())}

	app.Use(middleware.Logger())

	app.GET("/ping", func(ctx echo.Context) error {
		return ctx.JSON(http.StatusOK, map[string]string{
			"message": "pong",
			"status":  "successful start",
		})
	})

	app.HTTPErrorHandler = func(err error, c echo.Context) {
		code := http.StatusInternalServerError
		message := "Not found"

		if he, ok := err.(*echo.HTTPError); ok {
			code = he.Code
			message = he.Message.(string)
		}

		app.Logger.Error(err)
		c.JSON(code, map[string]interface{}{
			"status":  "false",
			"code":    code,
			"message": message,
		})
	}

	routes.UserRoutes(app)

	// Graceful quit
	c := make(chan os.Signal, 1)
	go func() {
		<-c
		database.DB.Close()
		_ = app.Shutdown(context.Background())
	}()
	signal.Notify(c, os.Interrupt)

	app.Logger.Fatal(app.Start(":" + os.Getenv("PORT")))
}
