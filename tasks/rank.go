package tasks

import "go-crud/cache"

// RestartDailyRank
func RestartDailyRank()error{

	return cache.RedisClient.Del("rank:daily").Err()
}