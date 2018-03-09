package healthcheck

import (
	"database/sql"
	"net/http"

	_ "github.com/lib/pq"
	"gopkg.in/labstack/echo.v3"
)

type HealthchecksCtrl struct {
	Config   map[string]interface{}
	Postgres interface{} `path:"/postgres" method:"GET"`
}

func (hc *HealthchecksCtrl) PostgresFunc(ctx echo.Context) error {
	db, err := sql.Open("postgres", hc.Config["PGConnStr"].(string))
	if err != nil {
		return ctx.String(http.StatusInternalServerError, "Connection error")
	}
	err = db.Ping()
	if err != nil {
		return ctx.String(http.StatusInternalServerError, "Connection error")
	}
	return ctx.String(http.StatusOK, "OK")
}
