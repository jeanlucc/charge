package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/theodo/scalab/config"
	"github.com/theodo/scalab/src/repositories"
	"github.com/theodo/scalab/src/security"
	"golang.org/x/crypto/bcrypt"
)

func SignIn(c echo.Context) error {
	cred := new(security.Credentials)
	if err := c.Bind(cred); err != nil {
		return c.String(http.StatusBadRequest, "Could not retrieve username password from body")
	}
	user := repositories.FindUserByEmail(cred.Username)

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(cred.Password)); err != nil {
		log.Println("login attempt with password mismatch on user: ", user.Id)
		return c.String(http.StatusUnauthorized, "")
	}

	sess, _ := session.Get(config.Cfg.Session.CookieName, c)
	sess.Values["email"] = user.Email
	sess.Save(c.Request(), c.Response())

	return c.String(http.StatusOK, "logged in user id: "+user.Id)
}

func Me(c echo.Context) error {
	sess, _ := session.Get(config.Cfg.Session.CookieName, c)
	return c.String(http.StatusOK, fmt.Sprintf("hello, here is your email: %v", sess.Values["email"]))
}

func SignUp(c echo.Context) error {
	return c.String(http.StatusOK, "NOT IMPLEMENTED")
}
