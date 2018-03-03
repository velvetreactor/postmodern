package healthcheck

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
	"gopkg.in/labstack/echo.v3"
)

type Healthchecker struct {
	ConnStr string
}

func New(connStr string) *Healthchecker {
	return &Healthchecker{ConnStr: connStr}
}

func (hc *Healthchecker) PostgresHealth(ctx echo.Context) error {
	db, err := sql.Open("postgres", hc.ConnStr)
	if err != nil {
		fmt.Println(err)
		return ctx.String(http.StatusInternalServerError, "Connection error")
	}
	err = db.Ping()
	if err != nil {
		fmt.Println(err)
		return ctx.String(http.StatusInternalServerError, "Connection error")
	}
	return ctx.String(http.StatusOK, "OK")
}
