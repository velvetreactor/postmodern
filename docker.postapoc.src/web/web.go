package web

import (
	"net/http"

	"gopkg.in/labstack/echo.v3"
)

type PagesCtrl struct {
	Config    map[string]interface{}
	Namespace string
	Home      interface{} `path:"/" method:"GET"`
}

func (pc *PagesCtrl) HomeFunc(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "Home")
}
