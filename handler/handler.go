package handler

import (
		"strconv"
		"net/http"
		"github.com/labstack/echo"
		"github.com/tkc/go-echo-server-sandbox/models/user"
)

type (
		resultStatus struct {
				Status bool
		}
		resultUserJson struct {
				User userModel.User
		}
		handler struct {
				userModel userModel.User
		}
		postUserData struct {
				Name   string `json:"name" form:"name"`
				Age    string `json:"age" form:"age"`
				AgeInt int
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

func (h handler) GetUser(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		user := h.userModel.Fetch(int64(id));
		return c.JSON(http.StatusOK, resultUserJson{User:user})
}

func (h handler) CreateUser(c echo.Context) error {
		post := new(postUserData)
		if err := c.Bind(post); err != nil {
				return err
		}
		post.changeAgeType()
		user := h.userModel.Create(post.Name, post.AgeInt);
		return c.JSON(http.StatusOK, user)
}

func (h handler) UpdateUser(c echo.Context) error {
		return c.JSON(http.StatusOK, "OK")
}

func (h handler) DeleteUser(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		h.userModel.Delete(int64(id));
		return c.JSON(http.StatusOK, resultStatus{Status:true})
}
