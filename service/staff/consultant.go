package staff

import (
	"fmt"
	"signin-go/common/types"
	"signin-go/global/time"
	"signin-go/internal/core"
	"signin-go/service/follow"

	"signin-go/repository/redis"
	"signin-go/repository/staff"
	"sync"
)

/*
获取门店的续卡相关的数据

	返回值为 /user_snapshot/accesstorenew 的 response
*/
func GetConsultantRenewData(ctx core.StdContext, studioID, staffUserID uint32) (result *types.AccessToRenewResponse, err error) {
	today := time.TodayDate()
	year, week := today.ISOWeek()
	weekKey := fmt.Sprintf("%v_%v", year, week)
	thisWeekStartDate := today.AddDate(0, 0, -int(today.Weekday()-1))
	thisWeekEndDate := thisWeekStartDate.AddDate(0, 0, 7)
	lastWeekStart := thisWeekStartDate.AddDate(0, 0, -7)
	lastWeekEnd := thisWeekStartDate
	nearly30StartDate := today.AddDate(0, 0, -30)
	nearly30EndDate := today.AddDate(0, 0, 1)

	var wg sync.WaitGroup
	consultantRenewData := []*types.ConsultantRenewData{}
	if staffUserID == 0 {
		idsStudioConsultants, err := staff.StudioConsultantOnlyID(ctx, studioID)
		if err == nil {
			mu := sync.Mutex{}
			appendDatas := func(data *types.ConsultantRenewData) {
				mu.Lock()
				defer func() {
					mu.Unlock()
				}()
				consultantRenewData = append(consultantRenewData, data)
			}

			for _, idsStudioConsultant := range idsStudioConsultants {
				wg.Add(1)
				go func(UserID uint32) {
					renewTargeValueRedisKey := redis.GetRenewTargeValueRedisKey(studioID, UserID)
					renewTargeValue, _ := redis.GetRenewTargeValue(ctx, renewTargeValueRedisKey)

					followUserStatus, _ := follow.GetConsultantFollowUserStatus(ctx, thisWeekStartDate, thisWeekEndDate, studioID, UserID)
					followUserIDs := make([]uint64, 0, len(followUserStatus))
					for followUserID := range followUserStatus {
						followUserIDs = append(followUserIDs, followUserID)
					}

					lastweekRenewRate, _ := follow.GetConsultantRenewRate(ctx, lastWeekStart, lastWeekEnd, studioID, UserID)
					nearly30RenewRate, _ := follow.GetConsultantRenewRate(ctx, nearly30StartDate, nearly30EndDate, studioID, UserID)
					appendDatas(&types.ConsultantRenewData{
						UserID:            UserID,
						TargetValue:       renewTargeValue[weekKey],
						LastweekRenewRate: lastweekRenewRate,
						Nearly30RenewRate: nearly30RenewRate,
						FollowUserIDs:     followUserIDs,
					})
					wg.Done()
				}(idsStudioConsultant.UserID)
			}
		}
	}

	followUserStatus, _ := follow.GetConsultantFollowUserStatus(ctx, thisWeekStartDate, thisWeekEndDate, studioID, staffUserID)
	followUserIDs := make([]uint64, 0, len(followUserStatus))
	for followUserID := range followUserStatus {
		followUserIDs = append(followUserIDs, followUserID)
	}

	lastweekRenewRate, _ := follow.GetConsultantRenewRate(ctx, lastWeekStart, lastWeekEnd, studioID, staffUserID)
	nearly30RenewRate, _ := follow.GetConsultantRenewRate(ctx, nearly30StartDate, nearly30EndDate, studioID, staffUserID)
	result = &types.AccessToRenewResponse{
		StudioID:          studioID,
		StaffUserID:       staffUserID,
		LastweekRenewRate: lastweekRenewRate,
		Nearly30RenewRate: nearly30RenewRate,
		FollowUserIDs:     followUserIDs,
	}
	wg.Wait()
	result.ConsultantTargets = consultantRenewData
	return
}
