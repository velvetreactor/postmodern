package web

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
	_ "github.com/lib/pq"
	"github.com/nycdavid/ziptie"
)

func getDbo(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", os.Getenv("PGCONN"))
	if err != nil {
		log.Print(err)
		t.Error("Cannot open PG connection")
	}
	err = db.Ping()
	if err != nil {
		log.Print(err)
		t.Error("Cannot open PG connection")
	}
	return db
}

func TestFetchTables(t *testing.T) {
	e := echo.New()
	ctrl := &TablesCtrl{Namespace: "/tables"}
	ziptie.Fasten(ctrl, e)
	cookieStore := sessions.NewCookieStore([]byte(os.Getenv("SESSION_SECRET")))
	e.Use(session.Middleware(cookieStore))

	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/tables", nil)
	ctx := e.NewContext(req, rec)
	ctx.Set("_session_store", cookieStore)
	sesn, _ := session.Get("session", ctx)
	sesn.Values["dbo"] = getDbo(t)
	populateDbForTest(sesn.Values["dbo"].(*sql.DB), t)
	e.ServeHTTP(rec, req)

	var tablesResp TablesResp
	json.NewDecoder(rec.Body).Decode(&tablesResp)
	if rec.Code != 200 {
		t.Error(fmt.Sprintf("Expected %d, got %d", 200, rec.Code))
	}
	if len(tablesResp.Tables) != 3 {
		t.Error(fmt.Sprintf("Expected %d tables, got %d", 3, len(tablesResp.Tables)))
	}
}
