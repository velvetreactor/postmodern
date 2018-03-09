package healthcheck

import (
	"fmt"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/nycdavid/ziptie"
	"gopkg.in/labstack/echo.v3"
)

func TestPostgresHealthError(t *testing.T) {
	e := echo.New()
	badConnStr := "xyznrbaaaa"
	ctrl := &HealthchecksCtrl{
		Config: map[string]interface{}{
			"PGConnStr": badConnStr,
		},
	}
	ziptie.Fasten(ctrl, e)

	rec := httptest.NewRecorder()
	req := httptest.NewRequest(echo.GET, "/healthchecks/postgres", nil)
	e.ServeHTTP(rec, req)

	if rec.Code != 500 {
		t.Fatal(fmt.Sprintf("Expecting status %s, got %d", "500", rec.Code))
	}
}

func TestPostgresHealthSuccess(t *testing.T) {
	e := echo.New()
	connStr := os.Getenv("PGCONN")
	ctrl := &HealthchecksCtrl{
		Config: map[string]interface{}{
			"PGConnStr": connStr,
		},
	}
	ziptie.Fasten(ctrl, e)

	rec := httptest.NewRecorder()
	req := httptest.NewRequest(echo.GET, "/healthchecks/postgres", nil)
	e.ServeHTTP(rec, req)

	if rec.Code != 200 {
		t.Fatal(fmt.Sprintf("Expecting status %s, got %d", "200", rec.Code))
	}
}
