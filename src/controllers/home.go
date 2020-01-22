package controllers

import (
	"github.com/labstack/echo/v4"
)

func Home(c echo.Context) error {
	return c.Render(200, "home.html", new(interface{}))
}
