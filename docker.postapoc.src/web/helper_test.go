package web

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
	_ "github.com/lib/pq"
	"github.com/velvetreactor/postapocalypse/testhelper"
)

func TestMain(m *testing.M) {
	dbo, err := sql.Open("postgres", os.Getenv("PGCONN"))
	if err != nil {
		log.Print(err)
		panic(err)
	}
	testhelper.CreateTestTables(dbo)
	testhelper.SeedDb(dbo)
	code := m.Run()
	teardownDb(dbo)
	os.Exit(code)
}

func teardownDb(dbo *sql.DB) {
	_, err := dbo.Exec("TRUNCATE items")
	if err != nil {
		log.Print(err)
		panic(err)
	}
}

func getDbo(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", os.Getenv("PGCONN"))
	if err != nil {
		log.Print(err)
		t.Error("Cannot open PG connection")
	}
	err = db.Ping()
	if err != nil {
		log.Print(err)
		t.Error("Cannot open PG connection")
	}
	return db
}

func setupSessionStore(e *echo.Echo) *sessions.CookieStore {
	cookieStore := sessions.NewCookieStore([]byte(os.Getenv("SESSION_SECRET")))
	e.Use(session.Middleware(cookieStore))
	return cookieStore
}
