package controllers

import "github.com/labstack/echo"

type homeDataResponse struct {
	Name string
}

func Home(c echo.Context) error {
	return c.Render(200, "home.html", homeDataResponse{Name: "JLC"})
}
