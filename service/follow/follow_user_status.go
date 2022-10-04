package follow

import (
	"signin-go/common/types"
	"signin-go/global/redis"
	"signin-go/global/time"
	"signin-go/internal/core"
	"signin-go/repository/follow"
	redisRepo "signin-go/repository/redis"
	"signin-go/repository/user_snapshot"
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
				syncMap.Store(followUserID, true)
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

/*
 */
func GetConsultantRenewRate(ctx core.StdContext, startTime, endTime time.Time, studioID, staffUserID uint32) (rate int, err error) {
	if studioID <= 0 {
		return
	}

	redisKey := redisRepo.GetConsultantRenewRateRedisKey(startTime, endTime, studioID, staffUserID)
	rate, err = redisRepo.GetConsultantRenewRate(ctx, redisKey)
	if err == nil {
		return
	}

	userSnapshotDatas, _ := user_snapshot.GetUserSnapshotData(ctx, &user_snapshot.UserSnapshotFilter{CreateTimeGE: startTime, CreateTimeLT: endTime, StudioID: studioID})
	userIDs := []uint32{}
	for _, userSnapshotData := range userSnapshotDatas {
		userIDs = append(userIDs, userSnapshotData.UserID)
	}

	lastUserSnapshotIDs, err := user_snapshot.GetLastUserSnapshotIDs(
		ctx,
		&user_snapshot.LastUserSnapshotIDsFilter{
			Time:           startTime,
			StudioID:       studioID,
			ExcludeUserIDs: userIDs,
		},
	)

	var earlyUserSnapshotDatas []*user_snapshot.UserSnapshotData
	if len(lastUserSnapshotIDs) > 0 {
		earlyUserSnapshotDatas, err = user_snapshot.GetUserSnapshotData(ctx, &user_snapshot.UserSnapshotFilter{IncludeIDs: lastUserSnapshotIDs})
	}
	for _, earlyUserSnapshotData := range earlyUserSnapshotDatas {
		userIDs = append(userIDs, earlyUserSnapshotData.UserID)
	}

	var userConsultantIDMap = follow.UserConsultantIDMap{}
	if staffUserID > 0 {
		userConsultantIDMap, _ = follow.GetUserLastFollowConsultantIDMap(ctx, userIDs)
	}

	allUserID := map[uint32]bool{}
	renewUserID := map[uint32]bool{}
	for _, userSnapshotData := range userSnapshotDatas {
		if !(staffUserID == userConsultantIDMap[userSnapshotData.UserID] && userSnapshotData.MembershipRemains+userSnapshotData.CouponRemains <= 35) {
			continue
		}
		allUserID[userSnapshotData.UserID] = true
		if userSnapshotData.RenewMembershipID > 0 {
			renewUserID[userSnapshotData.UserID] = true
		}
	}
	for _, earlyUserSnapshotData := range earlyUserSnapshotDatas {
		if !(staffUserID == userConsultantIDMap[earlyUserSnapshotData.UserID] && earlyUserSnapshotData.MembershipRemains+earlyUserSnapshotData.CouponRemains <= 35) {
			continue
		}
		allUserID[earlyUserSnapshotData.UserID] = true
		if earlyUserSnapshotData.RenewMembershipID > 0 {
			renewUserID[earlyUserSnapshotData.UserID] = true
		}
	}

	if len(allUserID) > 0 {
		rate = int(float64(len(renewUserID)) / float64(len(allUserID)) * 100)
		redis.Redis.Set(ctx, redisKey, rate, 1*24*time.Hour)
	}
	return
}
