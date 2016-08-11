package main

import (
		"fmt"
		"net/http"
		"github.com/labstack/echo"
		"github.com/labstack/echo/engine/standard"
		"github.com/tkc/go-echo-server-sandbox/handler"
		"github.com/tkc/go-echo-server-sandbox/models/user"
)

func status(c echo.Context) error {
		return c.String(http.StatusOK, "ok")
}

func main() {
		fmt.Println("sever start")
		e := echo.New()
		u := userModel.User{}
		h := handler.CreateHandler(u)
		e.GET("/", status)
		e.GET("/user/:id", h.GetUser)
		e.POST("/user", h.CreateUser)
		e.DELETE("/user", h.DeleteUser)
		e.Run(standard.New(":3001"))
}

