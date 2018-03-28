package web

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	queriesCtrl = &QueriesCtrl{Namespace: "/queries"}
	goodJson    = bytes.NewReader([]byte(`{ "query": "SELECT * FROM items WHERE items.name = 'Pencil';" }`))
	badQuery    = bytes.NewReader([]byte(`{ "query": "SELECT * FROM items JOIN belongings ON belonging_id = id;" }`))
)

func TestQueriesCreateNoAuthReturns401(t *testing.T) {
	e, _ := echoInit(queriesCtrl)

	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/queries", goodJson)

	e.ServeHTTP(rec, req)

	if rec.Code != 401 {
		t.Error(fmt.Sprintf("Expected status code %d, got %d", 401, rec.Code))
	}
}

func TestQueriesCreateGoodReqReturns200(t *testing.T) {
	e, cookieStore := echoInit(queriesCtrl)

	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/queries", goodJson)
	ctx := e.NewContext(req, rec)
	err := authenticateContext(ctx, cookieStore)
	if err != nil {
		t.Error("Error authenticating context:", err)
	}

	e.ServeHTTP(rec, req)

	if rec.Code != 200 {
		t.Error(fmt.Sprintf("Expected status code %d, got %d", 200, rec.Code))
	}
	var trs TableRows
	json.NewDecoder(rec.Body).Decode(&trs)
	firstRow := trs.Rows[0]
	pencilItem := map[string]string{
		"active": "true",
		"id":     "1",
		"name":   "Pencil",
	}
	for k, v := range pencilItem {
		if firstRow[k] != v {
			t.Error(fmt.Sprintf("Expected %s attribute to be %s, but got %s", k, v, firstRow[k]))
		}
	}
}

func TestQueriesCreateHandlesInvalidSql(t *testing.T) {
	e, cookieStore := echoInit(queriesCtrl)

	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/queries", badQuery)
	ctx := e.NewContext(req, rec)
	err := authenticateContext(ctx, cookieStore)
	if err != nil {
		t.Error("Error authenticating context:", err)
	}

	e.ServeHTTP(rec, req)
	errMsg := rec.Body.String()

	if rec.Code != 400 {
		t.Error(fmt.Sprintf("Expected status code %d, got %d", 400, rec.Code))
	}
	if errMsg == "" {
		t.Error("Expected non-blank error message, got:", errMsg)
	}
}

func TestQueriesCreateLimitsResultSet(t *testing.T) {
	selectAll := bytes.NewReader([]byte(`{ "query": "SELECT * FROM items;" }`))
	e, cookieStore := echoInit(queriesCtrl)

	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/queries", selectAll)
	ctx := e.NewContext(req, rec)
	err := authenticateContext(ctx, cookieStore)
	if err != nil {
		t.Error("Error authenticating context:", err)
	}

	e.ServeHTTP(rec, req)

	var trs TableRows
	json.NewDecoder(rec.Body).Decode(&trs)

	if len(trs.Rows) > 50 {
		t.Error(fmt.Sprintf("Expected rows to be capped at 50, got %d.", len(trs.Rows)))
	}
}

func TestQueriesCreateSupportUserDefinedLimit(t *testing.T) {
	queryWithLimit := bytes.NewReader([]byte(`{ "query": "SELECT * FROM items LIMIT 10;" }`))
	e, cookieStore := echoInit(queriesCtrl)

	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/queries", queryWithLimit)
	ctx := e.NewContext(req, rec)
	err := authenticateContext(ctx, cookieStore)
	if err != nil {
		t.Error("Error authenticating context:", err)
	}

	e.ServeHTTP(rec, req)

	var trs TableRows
	json.NewDecoder(rec.Body).Decode(&trs)

	if len(trs.Rows) != 10 {
		t.Error(fmt.Sprintf("Expected %d rows, got %d.", 10, len(trs.Rows)))
	}
}
