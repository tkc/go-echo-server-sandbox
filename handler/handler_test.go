package handler

import (
		"net/http"
		"net/http/httptest"
		"strings"
		"testing"
		"github.com/labstack/echo"
		"github.com/tkc/go-echo-server-sandbox/models/user"
		"github.com/labstack/echo/engine/standard"
		"github.com/stretchr/testify/assert"
)

var (
		userJSON = `{"Name":"post name","Age":"27"}`
)

func TestGetUser(t *testing.T) {
		e := echo.New()
		req := new(http.Request)
		rec := httptest.NewRecorder()
		c := e.NewContext(standard.NewRequest(req, e.Logger()), standard.NewResponse(rec, e.Logger()))
		c.SetPath("/users/:id")
		c.SetParamNames("id")
		c.SetParamValues("10")

		u := userModel.User{}
		h := CreateHandler(u)

		if assert.NoError(t, h.GetUser(c)) {
				assert.Equal(t, http.StatusOK, rec.Code)
				//t.Log(rec.Body.String())
		}
}

func TestCreateUser(t *testing.T) {

		e := echo.New()
		req, err := http.NewRequest(echo.POST, "/user", strings.NewReader(userJSON))
		if assert.NoError(t, err) {
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				rec := httptest.NewRecorder()
				c := e.NewContext(standard.NewRequest(req, e.Logger()), standard.NewResponse(rec, e.Logger()))

				u := userModel.User{}
				h := CreateHandler(u)

				if assert.NoError(t, h.CreateUser(c)) {
						assert.Equal(t, http.StatusOK, rec.Code)
						//t.Log(rec.Body.String())
				}
		}
}

