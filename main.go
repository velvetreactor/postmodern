package main

import (
	"database/sql"
	"net/http"

	"github.com/lib/pq"
	"gopkg.in/labstack/echo.v3"
)

func Home(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "Hello, world!")
}

func main() {
	e := echo.New()
	e.GET("/", Home)

	e.Logger.Fatal(e.Start(":1323"))
}
