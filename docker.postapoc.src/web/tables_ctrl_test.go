package web

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
	_ "github.com/lib/pq"
	"github.com/nycdavid/ziptie"
	"github.com/satori/go.uuid"
)

func TestFetchTables(t *testing.T) {
	e := echo.New()
	cookieStore := setupSessionStore(e)
	ctrl := &TablesCtrl{Namespace: "/tables"}
	ziptie.Fasten(ctrl, e)

	sampleUuid := uuid.NewV4()
	dbo := getDbo(t)
	populateDbForTest(dbo, t)
	DBObjects[sampleUuid] = dbo

	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/tables", nil)
	ctx := e.NewContext(req, rec)
	ctx.Set("_session_store", cookieStore)
	sesn, _ := session.Get("session", ctx)
	sesn.Values["uuid"] = sampleUuid.String()
	sesn.Save(ctx.Request(), ctx.Response())

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
