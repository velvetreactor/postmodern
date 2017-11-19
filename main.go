package main

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
	_ "github.com/lib/pq"
)

const (
	connStr = "user=t13 dbname=t13_web_dev sslmode=disable"
)

type DBCol string

func (dbc DBCol) Value() (driver.Value, error) {
	return "This is a dummy val", nil
}

func (dbc *DBCol) Scan(value interface{}) error {
	switch value.(type) {
	case []uint8:
		uintslc := value.([]uint8)
		*dbc = DBCol(string(uintslc))
	case time.Time:
		timeAttr := value.(time.Time)
		*dbc = DBCol(timeAttr.String())
	case string:
		stringAttr := value.(string)
		*dbc = DBCol(stringAttr)
	case int32:
		intAttr := value.(int)
		*dbc = DBCol(strconv.Itoa(intAttr))
	case bool:
		boolAttr := value.(bool)
		*dbc = DBCol(strconv.FormatBool(boolAttr))
	}
	return nil
}

type MyRenderer struct {
	templates *template.Template
}

func (r *MyRenderer) Render(w io.Writer, templateName string, data interface{}, ctx echo.Context) error {
	return r.templates.ExecuteTemplate(w, templateName, data)
}

func Sessions(ctx echo.Context) error {
	sess, e := session.Get("session", ctx)
	if e != nil {
		fmt.Println(e)
	}
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
	}
	// if pg auth is successful
	sess.Values["dbo"] = "db conn"
	sess.Save(ctx.Request(), ctx.Response())
	return ctx.String(http.StatusOK, "OK")
}

func Home(ctx echo.Context) error {
	sess, _ := session.Get("session", ctx)
	fmt.Println(sess.Values["id"])
	return ctx.Render(http.StatusOK, "home", "")
}

func GetTables(ctx echo.Context) error {
	db, err := sql.Open("postgres", connStr)
	var tableNames []string
	if err != nil {
		log.Fatal(err)
	}
	rows, err := db.Query("select table_name from information_schema.tables where table_schema = 'public' and table_type = 'BASE TABLE'")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var table_name string
		rows.Scan(&table_name)
		tableNames = append(tableNames, table_name)
	}
	return ctx.JSON(http.StatusOK, tableNames)
}

func GetRows(ctx echo.Context) error {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	tableName := ctx.QueryParam("table")
	if tableName == "" {
		log.Fatal("Table name required")
	}
	query := fmt.Sprintf("select * from %s limit 50", tableName)
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	rowMapSlc := generateRowMapSlc(rows)
	return ctx.JSON(http.StatusOK, rowMapSlc)
}

type ExecuteQueryReq struct {
	Query string `json:"query"`
}

func ExecuteQuery(ctx echo.Context) error {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	var eqr ExecuteQueryReq
	req := ctx.Request()
	dec := json.NewDecoder(req.Body)
	dec.Decode(&eqr)
	rows, err := db.Query(eqr.Query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	rowMapSlc := generateRowMapSlc(rows)
	return ctx.JSON(http.StatusOK, rowMapSlc)
}

func generateRowMapSlc(rows *sql.Rows) []map[string]DBCol {
	rowMapSlc := []map[string]DBCol{}
	cols, err := rows.Columns()
	if err != nil {
		log.Fatal(err)
	}
	results := make([]DBCol, len(cols))
	scanArgs := make([]interface{}, len(results))
	for i := range results {
		scanArgs[i] = &results[i]
	}
	for rows.Next() {
		rows.Scan(scanArgs...)
		rowMap := map[string]DBCol{}
		for idx, result := range results {
			rowMap[cols[idx]] = result
		}
		rowMapSlc = append(rowMapSlc, rowMap)
	}
	return rowMapSlc
}

func main() {
	e := echo.New()
	store := sessions.NewCookieStore([]byte("secret"))
	e.Use(session.Middleware(store))
	templates := template.Must(template.ParseGlob("src/templates/*.html"))
	e.Renderer = &MyRenderer{
		templates: templates,
	}
	e.Static("/dist", "dist")
	e.GET("/", Home)
	e.GET("/sessions", Sessions)
	e.POST("/query", ExecuteQuery)
	e.GET("/tables", GetTables)
	e.GET("/rows", GetRows)
	port := os.Getenv("PORT")
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}
