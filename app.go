package main

import (
	"net/http"
	"github.com/labstack/echo"
	"golang.org/x/crypto/acme/autocert"
	"github.com/labstack/echo/middleware"
	"github.com/go-echo-server-sandbox/config"
	"github.com/tkc/go-echo-server-sandbox/handler"
	"github.com/tkc/go-echo-server-sandbox/template"
	"github.com/tkc/go-echo-server-sandbox/models/user"
)

func status(c echo.Context) error {
	return c.String(http.StatusOK, "status ok")
}

func main() {

	e := echo.New()
	u := userModel.User{}
	h := handler.CreateHandler(u)

	template := template.GetTemplate()
	e.Renderer = &template

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	e.Static("/css", "./public/css")
	e.Static("/dist", "./public/dist")

	e.GET("/public", status)
	e.GET("/", status)
	e.GET("template", h.GetTemplate)
	e.GET("/user/:id", h.GetUser)
	e.POST("/user", h.CreateUser)
	e.DELETE("/user", h.DeleteUser)

	if config.IsProd() {
		e.AutoTLSManager.Cache = autocert.DirCache("./.cache")
		e.Pre(middleware.HTTPSRedirect())
		go func(c *echo.Echo) {
			e.Logger.Fatal(e.Start(":80"))
		}(e)
		e.Logger.Fatal(e.StartAutoTLS(":443"))
	} else {
		e.Logger.Fatal(e.Start(config.GetPort()))
	}
}
