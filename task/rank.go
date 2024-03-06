package task

import (
	"context"

	"github.com/ChenMiaoQiu/go-cloud-disk/cache"
)

// RestartDailyRank recalculate daily rank
func RestartDailyRank() error {
	return cache.RedisClient.Del(context.Background(), cache.DailyRankKey).Err()
}
