package web

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
	_ "github.com/lib/pq"
	"github.com/nycdavid/ziptie"
	"github.com/satori/go.uuid"
)

var (
	tablesCtrl = &TablesCtrl{Namespace: "/tables"}
)

func echoInit(ctrl ziptie.Ctrl) (*echo.Echo, *sessions.CookieStore) {
	e := echo.New()
	ziptie.Fasten(ctrl, e)
	return e, setupSessionStore(e)
}

func TestTablesIndex(t *testing.T) {
	e, cookieStore := echoInit(tablesCtrl)

	sampleUuid := uuid.NewV4()
	dbo := getDbo(t)
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

// Tables Show action
func TestTablesShowNonexistentTableReturns404(t *testing.T) {
	e, cookieStore := echoInit(tablesCtrl)

	sampleUuid := uuid.NewV4()
	dbo := getDbo(t)
	DBObjects[sampleUuid] = dbo

	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/tables/idontexist", nil)
	ctx := e.NewContext(req, rec)
	ctx.Set("_session_store", cookieStore)
	sesn, _ := session.Get("session", ctx)
	sesn.Values["uuid"] = sampleUuid.String()
	sesn.Save(ctx.Request(), ctx.Response())

	e.ServeHTTP(rec, req)

	if rec.Code != 404 {
		t.Error(fmt.Sprintf("Expected status code %d, got %d", 404, rec.Code))
	}
}

func TestTablesShowNoAuthReturns401(t *testing.T) {
	e, _ := echoInit(tablesCtrl)

	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/tables/items", nil)

	e.ServeHTTP(rec, req)

	if rec.Code != 401 {
		t.Error(fmt.Sprintf("Expected status code %d, got %d", 401, rec.Code))
	}
}

func TestTablesShowGoodReqReturns200(t *testing.T) {
	e, cookieStore := echoInit(tablesCtrl)

	sampleUuid := uuid.NewV4()
	dbo := getDbo(t)
	DBObjects[sampleUuid] = dbo

	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/tables/items", nil)
	ctx := e.NewContext(req, rec)
	ctx.Set("_session_store", cookieStore)
	sesn, _ := session.Get("session", ctx)
	sesn.Values["uuid"] = sampleUuid.String()
	sesn.Save(ctx.Request(), ctx.Response())

	e.ServeHTTP(rec, req)
	var resp TableRows
	json.NewDecoder(rec.Body).Decode(&resp)

	if rec.Code != 200 {
		t.Error(fmt.Sprintf("Expected status code %d, got %d", 200, rec.Code))
	}
	if len(resp.Rows) != 3 {
		t.Error(fmt.Sprintf("Expected %d rows, got %d", 3, len(resp.Rows)))
	}
	firstItem := resp.Rows[0]
	attrs := []map[string]string{
		{"id": "1"},
		{"name": "Pencil"},
		{"active": "true"},
	}
	for _, attr := range attrs {
		for k, v := range attr {
			if firstItem[k] != v {
				err := fmt.Sprintf("Expected %s for %s, but got %s", attr[k], k, firstItem[k])
				t.Error(err)
			}
		}
	}
}

func TestTablesShowEmptyTableReturnsArray(t *testing.T) {
	t.Skip("If a table is empty, it should return an empty erray and not null.")
}
