package web

import (
	"text/template"

	"github.com/labstack/echo/v4"
)

var templates = template.Must(template.ParseGlob("web/html/*.html"))

func IndexHandler(c echo.Context) error {
	return templates.ExecuteTemplate(c.Response().Writer, "index.html", nil)
}

func RegisterHandler(c echo.Context) error {
	return templates.ExecuteTemplate(c.Response().Writer, "register.html", nil)
}

func LoginHandler(c echo.Context) error {
	return templates.ExecuteTemplate(c.Response().Writer, "login.html", nil)
}

func SpeakHandler(c echo.Context) error {
	return templates.ExecuteTemplate(c.Response().Writer, "speak.html", nil)
}
