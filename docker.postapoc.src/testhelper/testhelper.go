package testhelper

import (
	"bufio"
	"database/sql"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/lib/pq"
)

func SeedDb(dbo *sql.DB, seedsPath string) {
	seeds, err := os.Open(seedsPath)
	if err != nil {
		log.Print(err)
		panic(err)
	}
	csvRdr := csv.NewReader(bufio.NewReader(seeds))
	rows, err := csvRdr.ReadAll()
	if err != nil {
		log.Print(err)
		panic(err)
	}
	for _, row := range rows {
		wrapQte := func(row []string) []string {
			var newRow []string
			for _, cell := range row {
				newRow = append(newRow, fmt.Sprintf("'%s'", cell))
			}
			return newRow
		}
		query := fmt.Sprintf("INSERT INTO items VALUES(%s)", strings.Join(wrapQte(row), ", "))
		_, err := dbo.Exec(query)
		if err != nil {
			log.Print(err)
			panic(err)
		}
	}
}

func CreateTestTables(dbo *sql.DB) {
	tables := []string{"items", "users", "posts"}
	attrs := "(id integer, name text, active boolean, other_id uuid, belonging_id integer)"
	for _, table := range tables {
		_, err := dbo.Exec(fmt.Sprintf("CREATE TABLE %s %s;", table, attrs))
		if err != nil {
			log.Print(err)
			panic(err)
		}
	}
	_, err := dbo.Exec("CREATE TABLE belongings (id integer)")
	if err != nil {
		log.Print(err)
		panic(err)
	}
}
