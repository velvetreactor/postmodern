package web

import (
	"database/sql"
	"fmt"
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
