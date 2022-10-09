package signin

import (
	"gsteps-go/global/redis"
	"gsteps-go/internal/core"
	"gsteps-go/repository/signin"
	"strconv"
	"time"
)

func GetAllSigninSpend(ctx core.StdContext, userID uint32) (spend float64) {
	var err error
	redisKey := signin.GetAllSigninSpendRedisKey(userID)
	redisData, _ := redis.Redis.Get(ctx, redisKey).Result()
	spend, err = strconv.ParseFloat(redisData, 64)
	if err == nil {
		return
	}
	spend, err = signin.GetAllSigninSpend(ctx, userID)
	if err == nil {
		redis.Redis.Set(ctx, redisKey, spend, 20*time.Hour)
	}
	return
}
