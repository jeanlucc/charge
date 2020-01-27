package router

import (
	"github.com/labstack/echo/v4"
	"github.com/theodo/scalab/src/controllers"
)

// Defines the routes of the application, it is called at server creation
func ConfRoutes(e *echo.Echo) {
	e.GET("/", controllers.Home)
	e.POST("/signin", controllers.SignIn)
	e.POST("/signup", controllers.SignUp)
	e.GET("/me", controllers.Me)
}
