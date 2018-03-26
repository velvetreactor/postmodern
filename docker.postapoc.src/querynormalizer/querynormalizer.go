package querynormalizer

import (
	"strings"
)

func Normalize(query string) string {
	if strings.Contains(query, ";") {
		newQuery := strings.Replace(query, ";", "", 1)
		return newQuery
	}
	return query
}

func HasLimit(query string) bool {
	lowercaseQuery := strings.ToLower(query)
	if strings.Contains(lowercaseQuery, "limit") {
		return true
	}
	return false
}
