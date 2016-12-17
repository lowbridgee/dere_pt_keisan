package main

import (
	"html/template"
	"io"
	"net/http"
	"os"
	"strconv"
    cal "testproject/calculate"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Template struct {
	templates *template.Template
}

type Config struct {
	point  string `json:"point"` //int64
	event string `json:"event"`//int64
	normal string `json:"normal"`//int64
}

type Result struct {
	Point int64
	Event int64
	Normal int64
	Normal_times int64
	Event_times int64
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
	e.GET("/", mainpage)
	e.POST("/",calculate)

	// Start server
	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}

func mainpage(c echo.Context) error {
	empty := Result{}
	return c.Render(http.StatusOK, "base", empty)
}

func calculate(c echo.Context) error {
	// Form情報取得
	point, _ := strconv.Atoi(c.FormValue("point"))
	event, _  := strconv.Atoi(c.FormValue("event"))
	normal, _ := strconv.Atoi(c.FormValue("normal"))

	// 計算準備
	p := cal.PlayStyle{Normal: int64(normal), Special: int64(event)}

	x, y := cal.Point2Time(int64(point), p)

	e := Result{Point: int64(point), Event: int64(event), Normal: int64(normal),Event_times: x, Normal_times: y}

	return c.Render(http.StatusOK, "base", e)
}
