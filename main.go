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
	"github.com/satori/go.uuid"
	"github.com/velvetreactor/postapocalypse/dbconnpool"
)

const (
	connStr = ""
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

type Session struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	DbName   string `json:"db_name"`
}

func Sessions(ctx echo.Context) error {
	var sessBody Session
	sess, err := session.Get("session", ctx)
	reqBody := ctx.Request().Body
	dec := json.NewDecoder(reqBody)
	dec.Decode(&sessBody)
	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		sessBody.Username,
		sessBody.Password,
		sessBody.Host,
		sessBody.Port,
		sessBody.DbName,
	)
	if err != nil {
		fmt.Println(err)
	}
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
	}
	k := sess.Values["uuid"].(string)
	if k == "" || dbconnpool.Connections[k] == nil {
		db, err := sql.Open("postgres", connStr)
		if err != nil {
			log.Fatal(err)
		}
		id := uuid.NewV4().String()
		dbconnpool.Connections[id] = db
		sess.Values["uuid"] = id
		sess.Save(ctx.Request(), ctx.Response())
	}
	resSlc, _ := json.Marshal(`{ code: 200 }`)
	return ctx.JSON(http.StatusOK, string(resSlc))
}

func Home(ctx echo.Context) error {
	return ctx.Render(http.StatusOK, "home", "")
}

func GetTables(ctx echo.Context) error {
	sess, _ := session.Get("session", ctx)
	id := sess.Values["uuid"].(string)
	db := dbconnpool.Connections[id]
	var tableNames []string
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
	secret := os.Getenv("STORE_SECRET")
	e.Use(session.Middleware(sessions.NewCookieStore([]byte(secret))))
	templates := template.Must(template.ParseGlob("src/templates/*.html"))
	e.Renderer = &MyRenderer{
		templates: templates,
	}
	e.Static("/dist", "dist")
	e.GET("/", Home)
	e.POST("/sessions", Sessions)
	e.POST("/query", ExecuteQuery)
	e.GET("/tables", GetTables)
	e.GET("/rows", GetRows)
	port := os.Getenv("PORT")
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}
