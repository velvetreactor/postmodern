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
	"github.com/satori/go.uuid"
)

var DBObjects = make(map[uuid.UUID]*sql.DB)

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
	storedSesnUuid, ok := sesn.Values["uuid"].(string)
	if !ok {
		log.Print("Uuid type assertion failed")
		return ctx.JSON(http.StatusUnauthorized, false)
	}
	sesnUuid, err := uuid.FromString(storedSesnUuid)
	if err != nil {
		log.Print(err)
		return ctx.JSON(http.StatusUnauthorized, false)
	}
	dbo := DBObjects[sesnUuid]
	if dbo.Ping() != nil {
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
	newUuid := uuid.NewV4()
	sesn.Values["uuid"] = newUuid.String()
	sesn.Save(ctx.Request(), ctx.Response())
	DBObjects[newUuid] = db
	return ctx.JSON(http.StatusOK, true)
}
