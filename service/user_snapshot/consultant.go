package user_snapshot

import (
	"gsteps-go/global/redis"
	"gsteps-go/internal/core"
	"gsteps-go/repository/follow"
	redisRepo "gsteps-go/repository/redis"
	"gsteps-go/repository/user_snapshot"
	"time"
)

/*
 */
func GetConsultantRenewRate(ctx core.StdContext, startTime, endTime time.Time, studioID, staffUserID uint32) (rate int) {
	if studioID <= 0 {
		return
	}

	redisKey := redisRepo.GetConsultantRenewRateRedisKey(startTime, endTime, studioID, staffUserID)
	rate, err := redisRepo.GetInt(ctx, redisKey)
	if err == nil {
		return
	}

	userSnapshotDatas, _ := user_snapshot.GetUserSnapshotData(ctx, &user_snapshot.UserSnapshotFilter{CreateTimeGE: startTime, CreateTimeLT: endTime, StudioID: studioID})
	userIDs := []uint32{}
	for _, userSnapshotData := range userSnapshotDatas {
		userIDs = append(userIDs, userSnapshotData.UserID)
	}

	lastUserSnapshotIDs, _ := user_snapshot.GetLastUserSnapshotIDs(
		ctx,
		&user_snapshot.LastUserSnapshotIDsFilter{
			Time:           startTime,
			StudioID:       studioID,
			ExcludeUserIDs: userIDs,
		},
	)

	var earlyUserSnapshotDatas []*user_snapshot.UserSnapshotData
	if len(lastUserSnapshotIDs) > 0 {
		earlyUserSnapshotDatas, _ = user_snapshot.GetUserSnapshotData(ctx, &user_snapshot.UserSnapshotFilter{IncludeIDs: lastUserSnapshotIDs})
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
