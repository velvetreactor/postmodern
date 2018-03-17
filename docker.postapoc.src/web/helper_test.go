package web

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
	_ "github.com/lib/pq"
)

func TestMain(m *testing.M) {
	dbo, err := sql.Open("postgres", os.Getenv("PGCONN"))
	if err != nil {
		log.Print(err)
		panic(err)
	}
	seedDb(dbo)
	code := m.Run()
	teardownDb(dbo)
	os.Exit(code)
}

func seedDb(dbo *sql.DB) {
	createTestTables(dbo)
	insertQrys := []string{
		"INSERT INTO items VALUES(1, 'Pencil', true)",
		"INSERT INTO items VALUES(2, 'Cup', false)",
		"INSERT INTO items VALUES(3, 'Lamp', true)",
	}
	for _, query := range insertQrys {
		_, err := dbo.Exec(query)
		if err != nil {
			log.Print(err)
			panic(err)
		}
	}
}

func teardownDb(dbo *sql.DB) {
	_, err := dbo.Exec("TRUNCATE items")
	if err != nil {
		log.Print(err)
		panic(err)
	}
}

func createTestTables(dbo *sql.DB) {
	tables := []string{"items", "users", "posts"}
	for _, table := range tables {
		_, err := dbo.Exec(fmt.Sprintf("CREATE TABLE %s (id integer, name text, active boolean);", table))
		if err != nil {
			log.Print(err)
			panic(err)
		}
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
