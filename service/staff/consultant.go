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

func GetConsultantRenewData(ctx core.StdContext, studioID, staffUserID uint32) (result *types.AccessToRenewResponse, err error) {
	result = &types.AccessToRenewResponse{
		StudioID:    studioID,
		StaffUserID: staffUserID,
	}

	today := time.TodayDate()
	year, week := today.ISOWeek()
	weekKey := fmt.Sprintf("%v_%v", year, week)
	thisWeekStartDate := today.AddDate(0, 0, -int(today.Weekday()-1))
	thisWeekEndDate := thisWeekStartDate.AddDate(0, 0, 7)
	// lastWeekStart := thisWeekStartDate.AddDate(0, 0, -7)
	// lastWeekEnd := thisWeekStartDate
	// nearly30StartDate := today.AddDate(0, 0, -30)
	// nearly30EndDate := today.AddDate(0, 0, 1)

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
				go func(staffUserID uint32) {
					renewTargeValueRedisKey := redis.GetRenewTargeValueRedisKey(studioID, staffUserID)
					renewTargeValue, _ := redis.GetRenewTargeValue(ctx, renewTargeValueRedisKey)

					followUserStatus, _ := follow.GetConsultantFollowUserStatus(ctx, thisWeekStartDate, thisWeekEndDate, studioID, staffUserID)
					followUserIDs := make([]uint64, 0, len(followUserStatus))
					for followUserID := range followUserStatus {
						followUserIDs = append(followUserIDs, followUserID)
					}

					appendDatas(&types.ConsultantRenewData{
						UserID:        staffUserID,
						TargetValue:   renewTargeValue[weekKey],
						FollowUserIDs: followUserIDs,
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

	result.FollowUserIDs = followUserIDs
	wg.Wait()
	result.ConsultantTargets = consultantRenewData
	return
}
