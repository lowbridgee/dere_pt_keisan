package main

import (
	"html/template"
	"io"
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Use template
	t := &Template{
		templates: template.Must(template.ParseGlob("public/views/*.html")),
	}
	e.Renderer = t

	// Use static file
	e.Static("/asset", "asset")

	// Route => handler
	e.GET("/", Mainpage)

	// Start server
	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}

func Mainpage(c echo.Context) error {
	return c.Render(http.StatusOK, "base", "")
}
