package redis

import (
	"fmt"
	"gsteps-go/global/redis"
	"gsteps-go/global/utils"
	"gsteps-go/internal/core"
	"gsteps-go/internal/errors"
	"strconv"

	"time"
)

/*
从redis中获取设置过的续卡目标值
*/
func GetRenewTargeValue(ctx core.StdContext, redisKey string) (data RenewTargeValue, err error) {
	data = RenewTargeValue{}

	redisData, err := redis.Redis.Get(ctx, redisKey).Result()
	if err != nil {
		return
	}

	err = utils.Json.Unmarshal([]byte(redisData), &data)
	return
}

/*
设置续卡目标值到redis
*/
func SetRenewTargeValue(ctx core.StdContext, redisKey string, redisData RenewTargeValue) {
	dataByte, err := utils.Json.Marshal(redisData)
	if err == nil {
		redis.Redis.Set(ctx, redisKey, string(dataByte), 1*365*24*time.Hour)
	}
}

/*
从redis获取 int 类型的值
*/
func GetInt(ctx core.StdContext, redisKey string) (data int, err error) {
	redisData, err := redis.Redis.Get(ctx, redisKey).Result()
	if err != nil {
		return
	}
	data, err = strconv.Atoi(redisData)
	return
}

/*
从redis获取 uint64 类型的值
*/
func GetUint64(ctx core.StdContext, redisKey string) (data uint64, err error) {
	redisData, err := redis.Redis.Get(ctx, redisKey).Result()
	if err != nil {
		return
	}
	data, err = strconv.ParseUint(redisData, 10, 64)
	return
}

/*
从redis中获取Uint64Slices数组

	List/Set/SortSet
	暂只支持Set
*/
func GetUint64Slice(ctx core.StdContext, redisKey string, redisType RedisType) (data []string, err error) {
	switch redisType {
	case List:
		err = errors.New("GetUint64Slice From Redis No【List】Func")
	case Set:
		data, err = redis.Redis.SMembers(ctx, redisKey).Result()
	case SortSet:
		err = errors.New("GetUint64Slice From Redis No【SortSet】Func")
	default:
		err = errors.New(fmt.Sprintf("GetUint64Slice From Redis No【%s】Func", redisType))
	}
	return
}
