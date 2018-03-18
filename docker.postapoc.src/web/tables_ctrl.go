package web

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
	_ "github.com/lib/pq"
	"github.com/satori/go.uuid"
)

type TablesResp struct {
	Tables []string `json:"tables"`
}

type TableRows struct {
	Rows []map[string]interface{} `json:"rows"`
}

type TablesCtrl struct {
	Namespace string
	Index     interface{} `path:"" method:"GET"`
	Show      interface{} `path:"/:tableName" method:"GET"`
}

func (ctrl *TablesCtrl) IndexFunc(ctx echo.Context) error {
	var tablesResp TablesResp
	sesn, _ := session.Get("session", ctx)
	uuidStr := sesn.Values["uuid"].(string)
	sesnUuid, err := uuid.FromString(uuidStr)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, false)
	}
	dbo := DBObjects[sesnUuid]

	rows, err := dbo.Query("SELECT table_name FROM information_schema.tables WHERE table_schema = 'public' AND table_type = 'BASE TABLE'")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	appendRows(rows, &tablesResp)

	return ctx.JSON(http.StatusOK, tablesResp)
}

func (ctrl *TablesCtrl) ShowFunc(ctx echo.Context) error {
	sesn, _ := session.Get("session", ctx)
	uuidStr, ok := sesn.Values["uuid"].(string)
	if !ok {
		return ctx.JSON(http.StatusUnauthorized, false)
	}
	sesnUuid, err := uuid.FromString(uuidStr)
	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, false)
	}
	dbo := DBObjects[sesnUuid]
	qry := fmt.Sprintf("SELECT * FROM %s LIMIT 50", ctx.Param("tableName"))
	rows, err := dbo.Query(qry)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, "Not found")
	}
	trs := &TableRows{}
	for rows.Next() {
		mapPGRowToTableRow(rows, trs) // at this point in time, represents the cursor at a specific row
	}
	return ctx.JSON(http.StatusOK, trs)
}

// Private
func appendRows(rows *sql.Rows, tablesResp *TablesResp) {
	for rows.Next() {
		var name string
		rows.Scan(&name)
		tablesResp.Tables = append(tablesResp.Tables, name)
	}
}

func mapPGRowToTableRow(rows *sql.Rows, trs *TableRows) {
	row := make(map[string]interface{})
	colTypes, err := rows.ColumnTypes()
	if err != nil {
		panic(err)
	}
	values := make([]interface{}, len(colTypes))
	for _, colType := range colTypes {
		value := new(interface{})
		row[colType.Name()] = value
		values = append(values, value)
	}
	startIdx := len(values) - len(colTypes)
	err = rows.Scan(values[startIdx:len(values)]...)
	if err != nil {
		log.Print(err)
	}
	for k, v := range row {
		colAssert(k, v, row)
	}
	trs.Rows = append(trs.Rows, row)
}

func colAssert(key string, value interface{}, row map[string]interface{}) {
	ptrIntfc := value.(*interface{})
	switch t := (*ptrIntfc).(type) {
	case int64:
		row[key] = strconv.Itoa(int(t))
	case bool:
		row[key] = strconv.FormatBool(t)
	}
}
