package web

import (
	"bytes"
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

func TestValidSessionCreationSendsCookie(t *testing.T) {
	e := echo.New()
	setupSessionStore(e)
	ctrl := &SessionsCtrl{Namespace: "/sessions"}
	ziptie.Fasten(ctrl, e)

	var jsonBody bytes.Buffer
	sesn := Session{ConnectionString: "postgres://postgres@postgres:5432/postgres?sslmode=disable"}
	json.NewEncoder(&jsonBody).Encode(sesn)
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/sessions", &jsonBody)

	e.ServeHTTP(rec, req)

	resCookie := rec.Header().Get("Set-Cookie")
	if resCookie == "" {
		t.Error("Expected Set-Cookie Header to be non-empty")
	}
}

func TestValidSessionCreationStoresDbo(t *testing.T) {
	e := echo.New()
	setupSessionStore(e)
	ctrl := &SessionsCtrl{Namespace: "/sessions"}
	ziptie.Fasten(ctrl, e)

	// Pre-authenticated request
	var jsonBody bytes.Buffer
	sesn := Session{ConnectionString: "postgres://postgres@postgres:5432/postgres?sslmode=disable"}
	json.NewEncoder(&jsonBody).Encode(sesn)
	req := httptest.NewRequest(http.MethodPost, "/sessions", &jsonBody)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)
	if rec.Code != 200 {
		t.Error(fmt.Sprintf("Expected %d, got %d", 200, rec.Code))
	}
	resp := rec.Result()
	respCookies := resp.Cookies()
	sesnCookie := respCookies[0]

	// Authenticated request
	req = httptest.NewRequest(http.MethodGet, "/sessions", nil)
	req.AddCookie(sesnCookie)
	rec = httptest.NewRecorder()

	e.ServeHTTP(rec, req)

	if rec.Code != 200 {
		t.Error(fmt.Sprintf("Expected %d, got %d", 200, rec.Code))
	}
}
