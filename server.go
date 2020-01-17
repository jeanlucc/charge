package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/theodo/scalab/config"
	"github.com/theodo/scalab/config/router"
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
	router.Routes(e)
	config.Templates(e)
	// Server
	e.Logger.Fatal(e.Start(":3000"))
}
