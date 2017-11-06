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
	rows, err := db.Query("select * from rings limit 10")
	if err != nil {
		log.Fatal(err)
	}
	cols, _ := rows.Columns()
	for _, colName := range cols {
		fmt.Println(colName)
	}
}
