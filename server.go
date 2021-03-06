package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/theodo/scalab/config"
	"github.com/theodo/scalab/config/router"
	"github.com/theodo/scalab/src/database"
)

func main() {
	e := echo.New()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	//CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))
	// Config
	config.LoadEnv()
	config.Configure()
	config.ConfTemplates(e)
	config.ConfSession(e)
	// Routing
	router.ConfRoutes(e)
	// Database
	database.Connect()
	// Server
	e.Logger.Fatal(e.Start(":3000"))
	// Shutdown
	database.Close()
}
