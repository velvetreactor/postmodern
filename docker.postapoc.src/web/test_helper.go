package web

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

func populateDbForTest(dbo *sql.DB, t *testing.T) {
	tables := []string{"items", "users", "posts"}
	for _, table := range tables {
		_, err := dbo.Exec(fmt.Sprintf("CREATE TABLE %s (id integer);", table))
		if err != nil {
			t.Error(err)
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
