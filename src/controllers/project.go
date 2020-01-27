package controllers

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/theodo/scalab/src/repositories"
)

type projectController struct{}

func NewProjectController() projectController {
	return projectController{}
}

func (_ *projectController) Get(c echo.Context) error {
	r := repositories.NewProjectRepository()
	projects, err := r.FindAll()
	if err != nil {
		log.Println(err)
	}
	return c.JSON(200, projects)
}
