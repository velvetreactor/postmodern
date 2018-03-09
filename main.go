package main

import (
	"os"

	"github.com/nycdavid/ziptie"
	"github.com/velvetreactor/postapocalypse/healthcheck"
	"gopkg.in/labstack/echo.v3"
)

func main() {
	e := echo.New()

	// Controllers
	healthchecksCtrl := &healthcheck.HealthchecksCtrl{
		Config: map[string]interface{}{
			"PGConnStr": os.Getenv("PGCONN"),
		},
	}

	ziptie.Fasten(healthchecksCtrl, e)

	// Start server
	e.Logger.Fatal(e.Start(":3000"))
}
