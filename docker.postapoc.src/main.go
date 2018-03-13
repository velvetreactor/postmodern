package main

import (
	"fmt"
	"os"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
	"github.com/nycdavid/ziptie"
	"github.com/velvetreactor/postapocalypse/web"
)

func main() {
	e := echo.New()

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
