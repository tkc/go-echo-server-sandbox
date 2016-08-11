package routes

import (
		"github.com/tkc/go-echo-server-sandbox/models"
		"github.com/labstack/echo"
		"net/http"
)

func Test(c echo.Context) error {
		user := model.User{}
		res := user.All()
		return c.JSON(http.StatusOK, res)
}
