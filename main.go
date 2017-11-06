package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"gopkg.in/labstack/echo.v3"
)

func Home(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "Hello, world!")
}

func main() {
	// e := echo.New()
	// e.GET("/", Home)
	//
	// e.Logger.Fatal(e.Start(":1323"))
	echo.New()
	connStr := "user=t13 dbname=t13_web_dev sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	rows, err := db.Query("select table_name from information_schema.tables where table_schema = 'public' and table_type = 'BASE TABLE'")
	defer rows.Close()
	for rows.Next() {
		var table_name string
		rows.Scan(&table_name)
		fmt.Println(table_name)
	}
}
