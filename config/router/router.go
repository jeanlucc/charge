package router

import (
	"github.com/labstack/echo/v4"
	"github.com/theodo/scalab/src/controllers"
	"github.com/theodo/scalab/src/middleware"
)

// Defines the routes of the application, it is called at server creation
func ConfRoutes(e *echo.Echo) {
	e.GET("/", controllers.Home)
	e.POST("/signin", controllers.SignIn)
	e.POST("/signup", controllers.SignUp)
	e.GET("/me", controllers.Me)
	apiRoutes(e)
}

func apiRoutes(e *echo.Echo) {
	apiG := e.Group("/api")
	apiG.Use(middleware.IsGranted(""))
	projectRoutes(apiG)
}

func projectRoutes(pg *echo.Group) {
	g := pg.Group("/projects")
	c := controllers.NewProjectController()
	g.GET("/", c.Get, middleware.IsGranted("ADMIN"))
	g.GET("/mine", c.GetMine)
}
