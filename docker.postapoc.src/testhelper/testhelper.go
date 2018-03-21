package testhelper

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func SeedDb(dbo *sql.DB) {
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

func CreateTestTables(dbo *sql.DB) {
	tables := []string{"items", "users", "posts"}
	for _, table := range tables {
		_, err := dbo.Exec(fmt.Sprintf("CREATE TABLE %s (id integer, name text, active boolean);", table))
		if err != nil {
			log.Print(err)
			panic(err)
		}
	}
}
