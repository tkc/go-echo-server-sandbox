package main

import (
	"net/http"
	"github.com/labstack/echo"
	"golang.org/x/crypto/acme/autocert"
	"github.com/labstack/echo/middleware"
	"github.com/tkc/go-echo-server-sandbox/config"
	"github.com/tkc/go-echo-server-sandbox/handler"
	"github.com/tkc/go-echo-server-sandbox/template"
	"github.com/tkc/go-echo-server-sandbox/models/user"
	"gopkg.in/go-playground/validator.v9"
)

type (
	CustomValidator struct {
		validator *validator.Validate
	}
)

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func status(c echo.Context) error {

	message := `
  ____    __
  / __/___/ /  ___
 / _// __/ _ \/ _ \
/___/\__/_//_/\___/ 

`
	return c.String(http.StatusOK, message)
}

func main() {

	e := echo.New()
	u := userModel.User{}
	h := handler.CreateHandler(u)

	template := template.GetTemplate()
	e.Renderer = &template
	e.Validator = &CustomValidator{validator: validator.New()}

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	e.Static("/css", "./public/css")
	e.Static("/dist", "./public/dist")

	e.GET("/", status)
	e.GET("template", h.GetTemplate)

	e.GET("/user/:id", h.Get)
	e.POST("/user", h.Create)
	e.PUT("/user", h.Update)
	e.DELETE("/user", h.Delete)

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
