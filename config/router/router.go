package router

import (
	"github.com/labstack/echo"
	"github.com/theodo/scalab/src/controllers"
)

// Defines the routes of the application, it is called at server creation
func Routes(e *echo.Echo) {
	e.GET("/", controllers.Home)
	e.POST("/login", controllers.Login)
}
