package redis

import (
	"fmt"
	"signin-go/global/redis"
	"signin-go/global/utils"
	"signin-go/internal/core"
	"time"
)

func GetRenewTargeValueRedisKey(studioID, staffUserID uint32) string {
	return fmt.Sprintf("HistoricalRenewTargeValue_%v_%v", studioID, staffUserID)
}

func GetRenewTargeValue(ctx core.Context, redisKey string) (data RenewTargeValue, err error) {
	data = RenewTargeValue{}

	redisData, err := redis.Redis.Get(ctx.RequestContext(), redisKey).Result()
	if err != nil {
		return
	}

	err = utils.Json.Unmarshal([]byte(redisData), &data)
	return
}

func SetRenewTargeValue(ctx core.Context, redisKey string, redisData RenewTargeValue) {
	dataByte, err := utils.Json.Marshal(redisData)
	if err == nil {
		redis.Redis.Set(ctx.RequestContext(), redisKey, string(dataByte), 1*365*24*time.Hour)
	}
}
