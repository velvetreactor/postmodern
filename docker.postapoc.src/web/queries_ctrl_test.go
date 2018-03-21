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

	e.ServeHTTP(rec, req)

	if err != nil {
		t.Error("Error authenticating context:", err)
	}
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
