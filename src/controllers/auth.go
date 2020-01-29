package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/theodo/scalab/src/repositories"

	"github.com/labstack/echo/v4"
	"github.com/theodo/scalab/src/security"
)

func SignIn(c echo.Context) error {
	cred := security.Credentials{}
	if err := c.Bind(&cred); err != nil {
		return c.String(http.StatusBadRequest, "Could not retrieve username password from body")
	}

	up := security.NewUserSigninProvider()
	user, err := up.SignIn(cred, c)
	if err != nil {
		if _, ok := err.(*security.PasswordMismatchError); ok {
			log.Println("login attempt with password mismatch on user: ", user.Id)
			return c.String(http.StatusUnauthorized, "")
		} else if _, ok := err.(*repositories.GetMappedResultError); ok {
			return c.String(http.StatusUnauthorized, "could not signin with provided credentials")
		} else {
			return c.JSON(http.StatusInternalServerError, "")
		}
	}

	return c.String(http.StatusOK, "logged in user id: "+user.Id)
}

func Me(c echo.Context) error {
	up := security.NewUserFromContextProvider()
	user, err := up.Get(c)
	if err != nil {
		return c.String(http.StatusUnauthorized, "")
	}

	return c.String(http.StatusOK, fmt.Sprintf("hello, here is your email: %v", user.Email))
}

func SignUp(c echo.Context) error {
	cred := security.ConfirmedCredentials{}
	if err := c.Bind(&cred); err != nil {
		return c.String(http.StatusBadRequest, "Could not retrieve username password from body")
	}

	uc := security.NewUserAccountCreator()
	user, err := uc.Create(cred)
	if err != nil {
		if _, ok := err.(*security.ConfirmedPasswordMismatchError); ok {
			return c.String(http.StatusBadRequest, "passwords mismatch")
		} else {
			return c.String(http.StatusInternalServerError, "Could not create account")
		}
	}

	return c.String(http.StatusOK, "Account created"+user.Id)
}
