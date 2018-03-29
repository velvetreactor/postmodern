package querynormalizer

import (
	"strconv"
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

func ToOffsetQty(page string) int {
	pageNum, _ := strconv.Atoi(page)
	if pageNum == 0 {
		pageNum = 1
	}
	offsetMultiplier := pageNum - 1
	offsetQty := 50 * offsetMultiplier
	return offsetQty
}
