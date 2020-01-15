package config

import (
	"html/template"
	"io"

	"github.com/labstack/echo"
)

const templatePath = "./templates/*.html"

type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}

// Defines the routes of the application, it is called at server creation
func Templates(e *echo.Echo) {
	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob(templatePath)),
	}
	e.Renderer = renderer
}
