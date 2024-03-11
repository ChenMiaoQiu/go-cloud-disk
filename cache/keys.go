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

// FileStoreKey use id to build file store info key in cache
func FileInfoStoreKey(id string) string {
	return fmt.Sprintf("file:cloud:%s", id)
}

// EmailConfirmKey use to store confirm code in cache
func EmailCodeKey(email string) string {
	return fmt.Sprintf("email:%s", email)
}

// RecentSendUserKey store user's required in recent
func RecentSendUserKey(email string) string {
	return fmt.Sprintf("user:confirm:%s", email)
}
