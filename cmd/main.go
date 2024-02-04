package main

import (
	"net/http"

	"github.com/CodeChefVIT/devsoc-backend-24/config"
	"github.com/labstack/echo/v4"
)

func init() {
	config.SanityCheck()
}

func main() {
	e := echo.New()
	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Pong")
	})
	e.Logger.Fatal(e.Start(":8080"))
}
