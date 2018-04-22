package handler

import (
	"net/http"
	"github.com/spf13/cast"
	"github.com/labstack/echo"
	"github.com/tkc/go-echo-server-sandbox/models/user"
)

type (
	handler struct {
		userModel userModel.User
	}
	resultJson struct {
		Id        int64
		Name      string
		Age       int
	}
	postUserData struct {
		Name   string `json:"Name" form:"name" validate:"required"`
		Age    int `json:"Age" form:"age" validate:"required"`
	}
	resultHtmlTemplate struct {
		Title string
		Users [] userModel.User
	}
)

func CreateHandler(u userModel.User) *handler {
	return &handler{u}
}

func (h handler) GetTemplate(c echo.Context) error {
	users := h.userModel.All(10, 0)
	res := resultHtmlTemplate{
		Title: "users result",
		Users: users,
	}
	return c.Render(http.StatusOK, "user", res)
}

func (h handler) Get(c echo.Context) error {
	user := h.userModel.Fetch(cast.ToInt64(c.Param("id")))
	return c.JSON(http.StatusOK, resultJson{
		Id:user.Id,
		Name:user.Name,
		Age:user.Age,
	})
}

func (h handler) Create(c echo.Context) error {
	post := new(postUserData)
	if err := c.Bind(post); err != nil {
		return c.JSON(http.StatusOK, err.Error())
	}
	if err := c.Validate(post); err != nil {
		return c.JSON(http.StatusOK, err.Error())
	}
	user := h.userModel.Create(post.Name, post.Age)
	return c.JSON(http.StatusOK, user.Id)
}

func (h handler) Update(c echo.Context) error {
	post := new(postUserData)
	if err := c.Bind(post); err != nil {
		return c.JSON(http.StatusOK, err.Error())
	}
	if err := c.Validate(post); err != nil {
		return c.JSON(http.StatusOK, err.Error())
	}
	user := h.userModel.Create(post.Name, post.Age)
	return c.JSON(http.StatusOK, user.Id)
}

func (h handler) Delete(c echo.Context) error {
	h.userModel.Delete(cast.ToInt64(c.Param("id")))
	return c.JSON(http.StatusOK, true)
}
