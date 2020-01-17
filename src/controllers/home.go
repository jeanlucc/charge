package controllers

import (
	"github.com/labstack/echo"
	"github.com/theodo/scalab/config"
)

type homeDataResponse struct {
	Name        string
	DatabaseUrl string
}

func Home(c echo.Context) error {
	cfg := config.Cfg
	return c.Render(200, "home.html", homeDataResponse{Name: "JLC", DatabaseUrl: cfg.Database.Url})
}
