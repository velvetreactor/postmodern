package main

import (
	"database/sql"
	"html/template"
	"io"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"gopkg.in/labstack/echo.v3"
)

const (
	connStr = "user=t13 dbname=t13_web_dev sslmode=disable"
)

type MyRenderer struct {
	templates *template.Template
}

func (r *MyRenderer) Render(w io.Writer, templateName string, data interface{}, ctx echo.Context) error {
	return r.templates.ExecuteTemplate(w, templateName, data)
}

func Home(ctx echo.Context) error {
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
	return ctx.Render(http.StatusOK, "home", tableNames)
}

func main() {
	e := echo.New()
	templates := template.Must(template.ParseGlob("src/templates/*.html"))
	e.Renderer = &MyRenderer{
		templates: templates,
	}
	e.GET("/", Home)

	e.Logger.Fatal(e.Start(":1323"))

}
