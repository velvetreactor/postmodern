package web

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
	_ "github.com/lib/pq"
)

type Session struct {
	ConnectionString string `json:"connectionString"`
}

type SessionsCtrl struct {
	Config    map[string]interface{}
	Namespace string
	Show      interface{} `path:"" method:"GET"`
	Create    interface{} `path:"" method:"POST"`
}

func (ctrl *SessionsCtrl) ShowFunc(ctx echo.Context) error {
	sesn, _ := session.Get("session", ctx)
	sesn.Options = &sessions.Options{MaxAge: 3600}
	pgConnStr := sesn.Values["pgConnStr"]
	sesn.Save(ctx.Request(), ctx.Response())
	if pgConnStr == nil {
		return ctx.JSON(http.StatusUnauthorized, false)
	}
	return ctx.JSON(http.StatusOK, true)
}

func (ctrl *SessionsCtrl) CreateFunc(ctx echo.Context) error {
	var sesnBody Session
	json.NewDecoder(ctx.Request().Body).Decode(&sesnBody)
	db, err := sql.Open("postgres", sesnBody.ConnectionString)
	if err != nil {
		log.Print(err)
		return ctx.JSON(http.StatusUnauthorized, false)
	}
	err = db.Ping()
	if err != nil {
		log.Print(sesnBody.ConnectionString)
		log.Print(err)
		return ctx.JSON(http.StatusUnauthorized, false)
	}
	sesn, _ := session.Get("session", ctx)
	sesn.Options = &sessions.Options{MaxAge: 3600}
	sesn.Values["dbo"] = db
	sesn.Save(ctx.Request(), ctx.Response())
	return ctx.JSON(http.StatusOK, true)
}
