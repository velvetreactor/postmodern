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

// .ToOffsetQty
func TestToOffsetQty(t *testing.T) {
	expected := 100
	page := "3"

	offsetQty := ToOffsetQty(page)

	if offsetQty != expected {
		errMsg := fmt.Sprintf("Expected offset quantity to be %d, got %d", expected, offsetQty)
		t.Error(errMsg)
	}
}

func TestToOffsetQtyHandlesZeroPageParam(t *testing.T) {
	expected := 0
	page := "0"

	offsetQty := ToOffsetQty(page)

	if offsetQty != expected {
		errMsg := fmt.Sprintf("Expected offset quantity to be %d, got %d", expected, offsetQty)
		t.Error(errMsg)
	}
}

func TestToOffsetQtyHandlesNonIntPageParam(t *testing.T) {
	expected := 0
	page := "asdf"

	offsetQty := ToOffsetQty(page)

	if offsetQty != expected {
		errMsg := fmt.Sprintf("Expected offset quantity to be %d, got %d", expected, offsetQty)
		t.Error(errMsg)
	}
}
