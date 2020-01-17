package controllers

import (
	"log"

	"github.com/labstack/echo"
	"github.com/theodo/scalab/src/repositories"
)

func Login(c echo.Context) error {
	login := c.FormValue("login")
	password := c.FormValue("password")

	user := repositories.FindUserByEmail(login)
	if password != user.Password {
		log.Println("password mismatch", password, user.Password)

		return c.String(401, "")
	}
	return c.String(200, "logged in user id: "+user.Id)
}
