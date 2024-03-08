package task

import (
	"context"

	"github.com/ChenMiaoQiu/go-cloud-disk/cache"
)

// RestartDailyRank recalculate daily rank
func RestartDailyRank() error {
	// dailyrank is highly likely that a big key, use
	// unlink delete for enhance execute speed
	return cache.RedisClient.Unlink(context.Background(), cache.DailyRankKey).Err()
}
