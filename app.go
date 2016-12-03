package main

import (
		"fmt"
		"io"
		"net/http"
		"text/template"
		"github.com/labstack/echo"
		"github.com/labstack/echo/engine/standard"
		"github.com/tkc/go-echo-server-sandbox/handler"
		"github.com/tkc/go-echo-server-sandbox/models/user"
)

type Template struct {
		templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
		return t.templates.ExecuteTemplate(w, name, data)
}

func status(c echo.Context) error {
		return c.String(http.StatusOK, "ok")
}

func main() {

		fmt.Println("sever start")

		t := &Template{
				templates: template.Must(template.ParseGlob("public/views/*.html")),
		}

		e := echo.New()
		u := userModel.User{}
		h := handler.CreateHandler(u)

		e.SetRenderer(t)
		e.GET("/", status)
		e.GET("template", h.GetTemplate)
		e.GET("/user/:id", h.GetUser)
		e.POST("/user", h.CreateUser)
		e.DELETE("/user", h.DeleteUser)
		e.Run(standard.New(":3001"))
}
