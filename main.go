package main

import (
	"os"

	"github.com/velvetreactor/postapocalypse/healthcheck"
	"gopkg.in/labstack/echo.v3"
)

func main() {
	e := echo.New()
	pgConnStr := os.Getenv("PGCONN")
	hc := healthcheck.New(pgConnStr)
	e.GET("/healthcheck/postgres", hc.PostgresHealth)
	e.Logger.Fatal(e.Start(":3000"))
}
