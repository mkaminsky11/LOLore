package main

import (
    "github.com/labstack/echo"
    "github.com/labstack/echo/engine/standard"
    "html/template"
    "io"
    "net/http"
)

type Template struct {
    templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
    return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	t := &Template{
	    templates: template.Must(template.ParseGlob("public/*.html")),
	}

    e := echo.New()

    e.Static("/", "static")
    e.SetRenderer(t)
    e.GET("/", func(c echo.Context) error {
        return c.Render(http.StatusOK, "index","index")
    })
    e.Run(standard.New(":1323"))
}