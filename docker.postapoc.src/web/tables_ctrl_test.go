package web

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"sort"
	"testing"

	_ "github.com/lib/pq"
	"github.com/satori/go.uuid"
)

var (
	tablesCtrl = &TablesCtrl{Namespace: "/tables"}
)

func TestTablesIndex(t *testing.T) {
	e, cookieStore := echoInit(tablesCtrl)

	sampleUuid := uuid.NewV4()
	dbo := getDbo(t)
	DBObjects[sampleUuid] = dbo

	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/tables", nil)
	ctx := e.NewContext(req, rec)
	err := authenticateContext(ctx, cookieStore)
	if err != nil {
		t.Error("Error authenticating context:", err)
	}

	e.ServeHTTP(rec, req)

	var tablesResp TablesResp
	json.NewDecoder(rec.Body).Decode(&tablesResp)
	if rec.Code != 200 {
		t.Error(fmt.Sprintf("Expected %d, got %d", 200, rec.Code))
	}
	if len(tablesResp.Tables) != 4 {
		t.Error(fmt.Sprintf("Expected %d tables, got %d", 4, len(tablesResp.Tables)))
	}
}

func TestTablesIndexSortsTableNames(t *testing.T) {
	e, cookieStore := echoInit(tablesCtrl)

	sampleUuid := uuid.NewV4()
	dbo := getDbo(t)
	DBObjects[sampleUuid] = dbo

	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/tables", nil)
	ctx := e.NewContext(req, rec)
	err := authenticateContext(ctx, cookieStore)
	if err != nil {
		t.Error("Error authenticating context:", err)
	}

	e.ServeHTTP(rec, req)

	var tablesResp TablesResp
	json.NewDecoder(rec.Body).Decode(&tablesResp)

	if !sort.StringsAreSorted(tablesResp.Tables) {
		t.Error("Expected tables array to be sorted.")
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
	err := authenticateContext(ctx, cookieStore)
	if err != nil {
		t.Error("Error authenticating context:", err)
	}

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
	err := authenticateContext(ctx, cookieStore)
	if err != nil {
		t.Error("Error authenticating context:", err)
	}

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
	e, cookieStore := echoInit(tablesCtrl)

	sampleUuid := uuid.NewV4()
	dbo := getDbo(t)
	DBObjects[sampleUuid] = dbo

	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/tables/posts", nil)
	ctx := e.NewContext(req, rec)
	err := authenticateContext(ctx, cookieStore)
	if err != nil {
		t.Error("Error authenticating context:", err)
	}

	e.ServeHTTP(rec, req)
	var resp TableRows
	json.NewDecoder(rec.Body).Decode(&resp)
	rt := reflect.TypeOf(resp.Rows)

	if rec.Code != 200 {
		t.Error(fmt.Sprintf("Expected status code %d, got %d", 200, rec.Code))
	}
	if rt.Kind() != reflect.Slice {
		t.Error("Expected slice.")
	}
}

func TestTablesShowCorrectlyConvertUuid(t *testing.T) {
	e, cookieStore := echoInit(tablesCtrl)

	sampleUuid := uuid.NewV4()
	dbo := getDbo(t)
	DBObjects[sampleUuid] = dbo

	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/tables/items", nil)
	ctx := e.NewContext(req, rec)
	err := authenticateContext(ctx, cookieStore)
	if err != nil {
		t.Error("Error authenticating context:", err)
	}

	itemUuid := uuid.NewV4()
	_, err = dbo.Exec(fmt.Sprintf("UPDATE items SET other_id = '%s';", itemUuid.String()))
	if err != nil {
		t.Error("Error creating test row:", err)
	}

	e.ServeHTTP(rec, req)

	var resp TableRows
	json.NewDecoder(rec.Body).Decode(&resp)
	if resp.Rows[0]["other_id"] != itemUuid.String() {
		t.Error(fmt.Sprintf("Expected uuid to be %s, got %s", itemUuid, resp.Rows[0]["other_id"]))
	}
}
