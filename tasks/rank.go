package tasks

import "giligili/cache"

// RestartDailyRank 重启一天的排名
func RestartDailyRank() error {
	return cache.RedisClient.Del(cache.Ctx, "rank:daily").Err()
}
