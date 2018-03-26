package querynormalizer

import (
	"fmt"
	"testing"
)

func TestParsesOutSemiColons(t *testing.T) {
	query := "SELECT * FROM items;"
	expectedQuery := "SELECT * FROM items"

	newQuery := Normalize(query)

	if newQuery != expectedQuery {
		t.Error(fmt.Sprintf("Expected query to be: %s, got %s", expectedQuery, newQuery))
	}
}

func TestHasLimit(t *testing.T) {
	query := "SELECT * FROM items LIMIT 10;"

	hasLimit := HasLimit(query)

	if hasLimit != true {
		t.Error(fmt.Sprintf("Expected hasLimit to be %t, got %t", true, hasLimit))
	}
}
