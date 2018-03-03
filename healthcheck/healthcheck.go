package healthcheck

import (
	"database/sql"
	"net/http"

	_ "github.com/lib/pq"
	"gopkg.in/labstack/echo.v3"
)

type Healthchecker struct {
	ConnStr string
	Echo    *echo.Echo
}

func New(connStr string) *Healthchecker {
	e := echo.New()
	return &Healthchecker{ConnStr: connStr, Echo: e}
}

func (hc *Healthchecker) PostgresHealth(ctx echo.Context) error {
	db, err := sql.Open("postgres", hc.ConnStr)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, "Connection error")
	}
	err = db.Ping()
	if err != nil {
		return ctx.String(http.StatusInternalServerError, "Connection error")
	}
	return ctx.String(http.StatusOK, "OK")
}
