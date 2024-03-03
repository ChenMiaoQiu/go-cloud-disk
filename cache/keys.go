package cache

import "fmt"

const (
	// DailyRankKey the daily rank of view
	DailyRankKey = "rank:daily"
)

// ShareKey use id to build share key in cache
func ShareKey(id string) string {
	return fmt.Sprintf("share:%s", id)
}
