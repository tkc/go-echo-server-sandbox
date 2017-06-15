package handler

import (
	"strconv"
	"net/http"
	"github.com/labstack/echo"
	"github.com/tkc/go-echo-server-sandbox/models/user"
)

type (
	handler struct {
		userModel userModel.User
	}
	resultStatus struct {
		Status bool
	}
	resultUserJson struct {
		User userModel.User
	}
	postUserData struct {
		Name   string `json:"name" form:"name"`
		Age    string `json:"age" form:"age"`
		AgeInt int
	}
)

/*
	Template Type
*/

type (
	userTemplate struct {
		Title string
		Users [] userModel.User
	}
)

func (p *postUserData) changeAgeType() {
	var i int
	i, _ = strconv.Atoi(p.Age)
	p.AgeInt = i
}

func CreateHandler(u userModel.User) *handler {
	return &handler{u}
}

func (h handler) GetTemplate(c echo.Context) error {
	users := h.userModel.All(10, 0)
	//users := userModel.User.All()//TODO
	res := userTemplate{
		Title: "users result",
		Users: users,
	}
	return c.Render(http.StatusOK, "user", res)
}

func (h handler) GetUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	user := h.userModel.Fetch(int64(id))
	return c.JSON(http.StatusOK, resultUserJson{User: user})
}

func (h handler) CreateUser(c echo.Context) error {
	post := new(postUserData)
	if err := c.Bind(post); err != nil {
		return err
	}
	post.changeAgeType()
	user := h.userModel.Create(post.Name, post.AgeInt)
	return c.JSON(http.StatusOK, user)
}

func (h handler) DeleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	h.userModel.Delete(int64(id))
	return c.JSON(http.StatusOK, resultStatus{Status: true})
}
