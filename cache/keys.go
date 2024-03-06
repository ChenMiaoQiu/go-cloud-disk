package cache

import "fmt"

const (
	// DailyRankKey the daily rank of view
	DailyRankKey = "rank:daily"
	// EmptyShare a set to store empty share key
	EmptyShare = "share:empty"
)

// ShareKey use id to build share key in cache
func ShareKey(id string) string {
	return fmt.Sprintf("share:%s", id)
}

// ShareInfoKey use id to build share info in cache
func ShareInfoKey(id string) string {
	return fmt.Sprintf("info:share:%s", id)
}
