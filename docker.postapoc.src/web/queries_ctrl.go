package web

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
	"github.com/satori/go.uuid"
	"github.com/velvetreactor/postapocalypse/querynormalizer"
)

type QueriesCtrl struct {
	Namespace string
	Create    interface{} `path:"" method:"POST"`
}

type Query struct {
	String string `json:"query"`
	Offset int    `json:"offset"`
}

func (ctrl *QueriesCtrl) CreateFunc(ctx echo.Context) error {
	var query Query
	sesn, _ := session.Get("session", ctx)
	uuidStr, ok := sesn.Values["uuid"].(string)
	if !ok {
		return ctx.JSON(http.StatusUnauthorized, false)
	}
	uuid, err := uuid.FromString(uuidStr)
	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, false)
	}
	dbo := DBObjects[uuid]
	json.NewDecoder(ctx.Request().Body).Decode(&query)
	if !querynormalizer.HasLimit(query.String) {
		query.String = querynormalizer.Normalize(query.String)
		query.String = fmt.Sprintf("%s LIMIT 50", query.String)
	}
	if query.Offset != 0 {
		query.String = fmt.Sprintf("%s OFFSET %d", query.String, query.Offset)
	}
	rows, err := dbo.Query(query.String)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	trs := &TableRows{}

	for rows.Next() {
		mapPGRowToTableRow(rows, trs) // at this point in time, represents the cursor at a specific row
	}
	return ctx.JSON(http.StatusOK, trs)
}
