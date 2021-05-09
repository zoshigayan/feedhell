package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/zoshigayan/feedhell/config"
	"github.com/zoshigayan/feedhell/controllers"
	"github.com/zoshigayan/feedhell/db"
	"net/http"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// e.Static("/static", "assets")

	db.Init()

	// Routes
	e.GET("/", func(c echo.Context) error {
		return c.Redirect(http.StatusMovedPermanently, "/feeds")
	})
	controllers.FeedController{}.Init(e.Group("/feeds"))
	controllers.SyncController{}.Init(e.Group("/sync"))

	e.Renderer = config.Renderer()
	e.Debug = true

	// Start server
	e.Logger.Fatal(e.Start(":4649"))
}
