package web

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
	"github.com/nycdavid/ziptie"
)

func TestInvalidSession(t *testing.T) {
	e := echo.New()
	cookieStore := sessions.NewCookieStore([]byte(os.Getenv("SESSION_SECRET")))
	e.Use(session.Middleware(cookieStore))
	ctrl := &SessionsCtrl{Namespace: "/sessions"}
	ziptie.Fasten(ctrl, e)
	req := httptest.NewRequest(http.MethodGet, "/sessions", nil)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)

	if rec.Code != 401 {
		t.Error(fmt.Sprintf("Expected %d, but got %d", 401, rec.Code))
	}
}

func TestValidSession(t *testing.T) {
	e := echo.New()
	cookieStore := sessions.NewCookieStore([]byte(os.Getenv("SESSION_SECRET")))
	e.Use(session.Middleware(cookieStore))
	ctrl := &SessionsCtrl{Namespace: "/sessions"}
	ziptie.Fasten(ctrl, e)
	req := httptest.NewRequest(http.MethodGet, "/sessions", nil)
	rec := httptest.NewRecorder()

	ctx := e.NewContext(req, rec)
	ctx.Set("_session_store", cookieStore)
	sesn, _ := session.Get("session", ctx)
	sesn.Options = &sessions.Options{MaxAge: 3600}
	sesn.Values["pgConnStr"] = "postgres://user@dbhost:5432/dummydb?sslmode=disable"
	e.ServeHTTP(rec, req)

	if rec.Code != 200 {
		t.Error(fmt.Sprintf("Expected %d, but got %d", 200, rec.Code))
	}
}
