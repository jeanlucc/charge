package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/theodo/scalab/src/repositories"
	"github.com/theodo/scalab/src/security"
)

type projectController struct{}

func NewProjectController() projectController {
	return projectController{}
}

func (_ *projectController) Get(c echo.Context) error {
	r := repositories.NewProjectRepository()
	projects, err := r.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, "")
	}
	return c.JSON(http.StatusOK, projects)
}

func (_ *projectController) GetMine(c echo.Context) error {
	up := security.NewUserFromContextProvider()
	u, err := up.Get(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "")
	}
	r := repositories.NewProjectRepository()
	projects, err := r.FindByOwner(u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "")
	}
	return c.JSON(http.StatusOK, projects)
}
