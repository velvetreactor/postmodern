package healthcheck

import (
	"net/http/httptest"
	"os"
	"testing"

	"gopkg.in/labstack/echo.v3"
)

func TestPostgresHealthError(t *testing.T) {
	connStr := "badconnstring"
	e := echo.New()
	hc := New(connStr)
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(echo.GET, "/postgres", nil)
	ctx := e.NewContext(req, rec)
	hc.PostgresHealth(ctx)
	code := ctx.Response().Status

	if code == 200 {
		t.Fatal("Expecting non-200 status")
	}
}

func TestPostgresHealthSuccess(t *testing.T) {
	connStr := os.Getenv("PGCONN")
	e := echo.New()
	hc := New(connStr)
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(echo.GET, "/postgres", nil)
	ctx := e.NewContext(req, rec)
	hc.PostgresHealth(ctx)
	code := ctx.Response().Status

	if code != 200 {
		t.Fatal("Expecting non-200 status")
	}
}
