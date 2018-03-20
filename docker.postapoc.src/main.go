package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
	_ "github.com/lib/pq"
	"github.com/nycdavid/ziptie"
	"github.com/velvetreactor/postapocalypse/testhelper"
	"github.com/velvetreactor/postapocalypse/web"
)

func main() {
	e := echo.New()

	if os.Getenv("ENV") == "test" {
		RunDbSetup()
	}

	// Middleware
	e.Use(session.Middleware(sessions.NewCookieStore([]byte(os.Getenv("SESSION_SECRET")))))

	// Controllers
	pagesCtrl := &web.PagesCtrl{Namespace: ""}
	sessionsCtrl := &web.SessionsCtrl{Namespace: "/sessions"}
	tablesCtrl := &web.TablesCtrl{Namespace: "/tables"}
	ziptie.Fasten(pagesCtrl, e)
	ziptie.Fasten(sessionsCtrl, e)
	ziptie.Fasten(tablesCtrl, e)

	// Rendering
	renderer := web.NewRenderer("web/templates/*.html")
	e.Renderer = renderer

	// Static assets
	e.Static("/dist", "dist")
	e.Static("/static", "static")

	// Start server
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", os.Getenv("PORT"))))
}

func RunDbSetup() {
	dbo, err := sql.Open("postgres", os.Getenv("PGCONN"))
	if err != nil {
		log.Print(err)
	}
	err = dbo.Ping()
	if err != nil {
		log.Print(err)
	}
	// testhelper.CreateTestTables(dbo)
	testhelper.SeedDb(dbo)
}
