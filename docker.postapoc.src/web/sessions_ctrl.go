package web

import (
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
)

type SessionsCtrl struct {
	Config    map[string]interface{}
	Namespace string
	Show      interface{} `path:"" method:"GET"`
}

func (ctrl *SessionsCtrl) ShowFunc(ctx echo.Context) error {
	sesn, _ := session.Get("session", ctx)
	sesn.Options = &sessions.Options{MaxAge: 3600}
	pgConnStr := sesn.Values["pgConnStr"]
	sesn.Save(ctx.Request(), ctx.Response())
	if pgConnStr == nil {
		return ctx.JSON(http.StatusUnauthorized, false)
	}
	return ctx.JSON(http.StatusOK, true)
}
