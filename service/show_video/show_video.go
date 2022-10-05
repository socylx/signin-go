package show_video

import (
	"signin-go/global/redis"
	"signin-go/internal/core"
	"signin-go/repository/show_video"
	"strconv"
	"time"
)

func GetShowVideoCount(ctx core.StdContext, userID uint32) (count int64) {
	var err error

	redisKey := show_video.GetShowVideoCountRedisKey(userID)
	showVideoCountRedisData, _ := redis.Redis.Get(ctx, redisKey).Result()
	count, err = strconv.ParseInt(showVideoCountRedisData, 10, 64)
	if err == nil {
		return
	}
	count, err = show_video.GetShowVideoCount(ctx, userID)
	if err == nil {
		redis.Redis.Set(ctx, redisKey, count, 20*time.Hour)
	}
	return
}
