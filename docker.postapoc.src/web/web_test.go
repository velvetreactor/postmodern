package web

import (
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/nycdavid/ziptie"
	"gopkg.in/labstack/echo.v3"
)

func TestHomePage(t *testing.T) {
	e := echo.New()
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(echo.GET, "/", nil)
	pagesCtrl := &PagesCtrl{Namespace: ""}
	ziptie.Fasten(pagesCtrl, e)

	e.ServeHTTP(rec, req)

	if rec.Code != 200 {
		t.Error(fmt.Sprintf("Expected %d, got %d", 200, rec.Code))
	}
}
