package main

import (
		"testing"
		"net/http"
		"net/http/httptest"
		"github.com/labstack/echo"
		"github.com/labstack/echo/engine/standard"
)

func TestHandler(t *testing.T) {

		e := echo.New()
		req := new(http.Request)
		rec := httptest.NewRecorder()
		c := e.NewContext(standard.NewRequest(req, e.Logger()), standard.NewResponse(rec, e.Logger()))
		c.SetPath("/")

		//u := &userModelStub{name: "dummy tannai"}
		//h := handler.NewHandler()
		//h.GetUser(c)
		//if rec.Body.String() != "{\"id\":\"dummy tannai\"}" {
		//		t.Errorf("expected response id: dummy tannai, got %s", rec.Body.String())
		//}

		res := rec.Body.String()
		t.Log(res)
}