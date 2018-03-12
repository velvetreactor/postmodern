package web

import (
	"net/http"

	"github.com/labstack/echo"
)

type PagesCtrl struct {
	Config    map[string]interface{}
	Namespace string
	Home      interface{} `path:"/" method:"GET"`
}

func (pc *PagesCtrl) HomeFunc(ctx echo.Context) error {
	return ctx.Render(http.StatusOK, "home", nil)
}
