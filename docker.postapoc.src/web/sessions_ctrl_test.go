package web

import (
	"bytes"
	"database/sql"
	"encoding/json"
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

func setupSessionStore(e *echo.Echo) {
	cookieStore := sessions.NewCookieStore([]byte(os.Getenv("SESSION_SECRET")))
	e.Use(session.Middleware(cookieStore))
}

func TestInvalidSession(t *testing.T) {
	e := echo.New()
	setupSessionStore(e)
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
	setupSessionStore(e)
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

func TestValidSessionCreation(t *testing.T) {
	e := echo.New()
	cookieStore := sessions.NewCookieStore([]byte(os.Getenv("SESSION_SECRET")))
	e.Use(session.Middleware(cookieStore))
	ctrl := &SessionsCtrl{Namespace: "/sessions"}
	ziptie.Fasten(ctrl, e)
	rec := httptest.NewRecorder()
	var jsonBody bytes.Buffer
	sesn := Session{ConnectionString: "postgres://postgres@postgres:5432/postgres?sslmode=disable"}
	json.NewEncoder(&jsonBody).Encode(sesn)

	req := httptest.NewRequest(http.MethodPost, "/sessions", &jsonBody)
	ctx := e.NewContext(req, rec)
	ctx.Set("_session_store", cookieStore)
	storedSesn, _ := session.Get("session", ctx)
	e.ServeHTTP(rec, req)

	if rec.Code != 200 {
		t.Error(fmt.Sprintf("Expected %d, got %d", 200, rec.Code))
	}
	if storedSesn.Values["dbo"].(*sql.DB).Ping() != nil {
		t.Error("Database object unavailable")
	}
}
