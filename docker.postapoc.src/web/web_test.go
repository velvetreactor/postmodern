package web

import (
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/nycdavid/ziptie"
)

func TestHomePage(t *testing.T) {
	e := echo.New()
	renderer := NewRenderer("templates/*.html")
	e.Renderer = renderer
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(echo.GET, "/", nil)

	pagesCtrl := &PagesCtrl{Namespace: ""}
	ziptie.Fasten(pagesCtrl, e)

	e.ServeHTTP(rec, req)
	conType := rec.HeaderMap.Get("Content-Type")

	if rec.Code != 200 {
		t.Error(fmt.Sprintf("Expected: %d status code, got: %d", 200, rec.Code))
	}
	if conType != echo.MIMETextHTMLCharsetUTF8 {
		t.Error(fmt.Sprintf("Expected: %s content-type, got: %s", echo.MIMETextHTMLCharsetUTF8, conType))
	}
}
