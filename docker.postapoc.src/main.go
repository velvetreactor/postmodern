package main

import (
	"fmt"
	"os"

	"github.com/nycdavid/ziptie"
	"github.com/velvetreactor/postapocalypse/web"
	"gopkg.in/labstack/echo.v3"
)

func main() {
	e := echo.New()

	// Controllers
	pagesCtrl := &web.PagesCtrl{Namespace: ""}
	ziptie.Fasten(pagesCtrl, e)

	// Rendering
	renderer := web.NewRenderer("web/templates/*.html")
	e.Renderer = renderer

	// Static assets
	e.Static("/dist", "dist")
	e.Static("/static", "static")

	// Start server
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", os.Getenv("PORT"))))
}
