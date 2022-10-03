package user_snapshot

import (
	"fmt"
	"signin-go/global/time"
	"signin-go/internal/code"
	"signin-go/internal/core"
	"signin-go/internal/validation"
	"signin-go/repository/redis"
)

type accesstorenewRequest struct {
	StudioID    uint32 `form:"studio_id" binding:"required"`
	StaffUserID uint32 `form:"staff_user_id"`
	TargetValue uint32 `form:"target_value"`
	IsSet       uint32 `form:"is_set"`
}

func accesstorenew(c core.Context) {
	request := new(accesstorenewRequest)
	if err := c.ShouldBindForm(request); err != nil {
		c.AbortWithError(core.Error(
			code.ParamBindError,
			validation.Error(err)).WithError(err),
		)
		return
	}
	if request.StudioID <= 0 {
		c.AbortWithError(core.Error(code.ParamBindError, "无门店ID参数"))
		return
	}

	today := time.TodayDate()
	year, week := today.ISOWeek()
	weekKey := fmt.Sprintf("%v_%v", year, week)
	// thisWeekStartDate := today.AddDate(0, 0, -int(today.Weekday()-1))
	// thisWeekEndDate := thisWeekStartDate.AddDate(0, 0, 7)
	// lastWeekStart := thisWeekStartDate.AddDate(0, 0, -7)
	// lastWeekEnd := thisWeekStartDate
	// nearly30StartDate := today.AddDate(0, 0, -30)
	// nearly30EndDate := today.AddDate(0, 0, 1)

	renewTargeValueRedisKey := redis.GetRenewTargeValueRedisKey(request.StudioID, request.StaffUserID)
	renewTargeValue, _ := redis.GetRenewTargeValue(c.RequestContext(), renewTargeValueRedisKey)
	if request.IsSet != 0 { // 设置
		renewTargeValue[weekKey] = uint64(request.TargetValue)
		redis.SetRenewTargeValue(c.RequestContext(), renewTargeValueRedisKey, renewTargeValue)
	}

	// consultantRenewData :=

}
