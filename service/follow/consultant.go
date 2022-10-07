package follow

import (
	"signin-go/common/types"
	"signin-go/global/redis"
	"signin-go/global/time"
	"signin-go/internal/core"
	"signin-go/repository/follow"
	redisRepo "signin-go/repository/redis"
	"strconv"
	"sync"
)

/*
获取某段时间内某个门店有续卡跟进的用户ID
*/
func GetConsultantFollowUserStatus(ctx core.StdContext, startTime, endTime time.Time, studioID, staffUserID uint32) (followUserStatus types.FollowUserStatus, err error) {
	followUserStatus = types.FollowUserStatus{}

	var (
		syncMap sync.Map
		wg      sync.WaitGroup
		days    int
	)

	today := time.TodayDate()
	for {
		queryStartTime := startTime.AddDate(0, 0, days)
		queryEndTime := queryStartTime.AddDate(0, 0, 1)

		wg.Add(1)
		go func(sTime, eTime time.Time) {
			defer func() {
				wg.Done()
			}()

			redisKey := redisRepo.GetConsultantFollowUserIDsRedisKey(sTime, eTime, studioID, staffUserID)
			redisData, err := redisRepo.GetUint64Slice(ctx, redisKey, redisRepo.Set)
			if err == nil && len(redisData) > 0 {
				for _, item := range redisData {
					userID, err := strconv.ParseUint(item, 10, 64)
					if err != nil {
						continue
					}
					syncMap.Store(userID, true)
				}
				return
			}

			followUserIDs, err := follow.GetFollowUserIDs(ctx, sTime, eTime, studioID, staffUserID)
			if err != nil {
				return
			}

			for _, followUserID := range followUserIDs {
				if queryStartTime.Before(today) {
					redis.Redis.SAdd(ctx, redisKey, followUserID)
				}
				syncMap.Store(uint64(followUserID), true)
			}
		}(queryStartTime, queryEndTime)

		if queryEndTime.Equal(endTime) || queryEndTime.After(endTime) {
			break
		}
		days += 1
	}
	wg.Wait()

	syncMap.Range(func(k, v interface{}) bool {
		if k != "user_id" && k != "user_before_member_id" {
			followUserStatus[k.(uint64)] = v.(bool)
		}
		return true
	})
	return
}
