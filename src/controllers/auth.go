package controllers

import (
	"log"
	"net/http"

	"github.com/theodo/scalab/src/repositories"

	"github.com/labstack/echo/v4"
	"github.com/theodo/scalab/src/security"
)

func SignIn(c echo.Context) error {
	cred := security.Credentials{}
	if err := c.Bind(&cred); err != nil {
		return c.JSON(http.StatusBadRequest, message{"Could not retrieve username password from body"})
	}

	up := security.NewUserSigninProvider()
	user, err := up.SignIn(cred, c)
	if err != nil {
		if _, ok := err.(*security.PasswordMismatchError); ok {
			log.Println("login attempt with password mismatch on user: ", user.Id)
			return c.JSON(http.StatusUnauthorized, message{""})
		} else if _, ok := err.(*repositories.GetMappedResultError); ok {
			return c.JSON(http.StatusUnauthorized, message{"could not signin with provided credentials"})
		} else {
			return c.JSON(http.StatusInternalServerError, message{""})
		}
	}

	return c.JSON(http.StatusOK, "")
}

func Me(c echo.Context) error {
	up := security.NewUserFromContextProvider()
	user, err := up.Get(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, message{""})
	}

	return c.JSON(http.StatusOK, user)
}

func SignUp(c echo.Context) error {
	cred := security.ConfirmedCredentials{}
	if err := c.Bind(&cred); err != nil {
		return c.JSON(http.StatusBadRequest, message{"Could not retrieve username password from body"})
	}

	uc := security.NewUserAccountCreator()
	user, err := uc.Create(cred)
	if err != nil {
		if _, ok := err.(*security.ConfirmedPasswordMismatchError); ok {
			return c.JSON(http.StatusBadRequest, message{"passwords mismatch"})
		} else {
			return c.JSON(http.StatusInternalServerError, message{"Could not create account"})
		}
	}

	return c.JSON(http.StatusOK, user.Id)
}
