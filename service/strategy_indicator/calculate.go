package strategy_indicator

import (
	"fmt"
	"signin-go/global/time"
	"signin-go/global/utils"
	"signin-go/internal/errors"
	"signin-go/repository/coupon"
	"signin-go/repository/course_level"
	orderRepo "signin-go/repository/order"
	"signin-go/repository/source"
	"signin-go/repository/strategy"
	"signin-go/repository/users"
	"strconv"
)

var noCalculateScoreError = errors.New("无计算得分")

type CalculateFunc func(userData *users.Data, strategyIndicatorRules []*strategy.StrategyIndicatorRule) (score *users.Score, err error)

var strategyIndicatorCalculateFunc = map[string]CalculateFunc{
	"remain": func(userData *users.Data, strategyIndicatorRules []*strategy.StrategyIndicatorRule) (score *users.Score, err error) {
		var (
			remains   float32
			todayDate = time.TodayDate()
		)
		for _, membership := range userData.Memberships {
			remains += membership.Remains
		}
		for _, couponAlloc := range userData.CouponAllocData.CouponAllocs {
			deadline := couponAlloc.Deadline
			if deadline != time.TimeZeroTime || deadline.Before(todayDate) {
				continue
			}
			if couponAlloc.Coupon.Type == coupon.CashType && couponAlloc.Coupon.AmountType == coupon.NumberAmountType {
				remains += couponAlloc.Remains
			}
		}
		var remainsFloat64 float64 = float64(remains)
		for _, strategyIndicatorRule := range strategyIndicatorRules {
			min, err1 := strconv.ParseFloat(strategyIndicatorRule.Min, 64)
			max, err2 := strconv.ParseFloat(strategyIndicatorRule.Max, 64)
			if err1 != nil || err2 != nil {
				continue
			}
			if min <= remainsFloat64 && (remainsFloat64 < max || (min > 0 && max <= 0)) {
				score = &users.Score{
					ID:    strategyIndicatorRule.ID,
					Name:  strategyIndicatorRule.Name,
					Score: float64(strategyIndicatorRule.Score),
				}
				return
			}
		}
		return nil, noCalculateScoreError
	},
	"pin_class": func(userData *users.Data, strategyIndicatorRules []*strategy.StrategyIndicatorRule) (score *users.Score, err error) {
		var (
			spend     float32
			todayDate = time.TodayDate()
			startDate = todayDate.AddDate(0, 0, -30)
			endDate   = todayDate.AddDate(0, 0, 1)
		)
		for _, signin := range userData.Signins {
			activityStartTime := signin.ActivityStartTime
			if activityStartTime.After(startDate) && activityStartTime.Before(endDate) {
				spend += signin.SigninSpend
			}
		}
		var spendFloat64 float64 = float64(spend)
		for _, strategyIndicatorRule := range strategyIndicatorRules {
			min, err1 := strconv.ParseFloat(strategyIndicatorRule.Min, 64)
			max, err2 := strconv.ParseFloat(strategyIndicatorRule.Max, 64)
			if err1 != nil || err2 != nil {
				continue
			}
			if min <= spendFloat64 && (spendFloat64 < max || (min > 0 && max <= 0)) {
				score = &users.Score{}
				score.ID = strategyIndicatorRule.ID
				score.Name = strategyIndicatorRule.Name
				score.Score = float64(strategyIndicatorRule.Score)
				return
			}
		}
		return nil, noCalculateScoreError
	},
	"renew_count": func(userData *users.Data, strategyIndicatorRules []*strategy.StrategyIndicatorRule) (score *users.Score, err error) {
		renewCount := len(userData.Memberships) - 1
		for _, strategyIndicatorRule := range strategyIndicatorRules {
			min, err1 := strconv.Atoi(strategyIndicatorRule.Min)
			max, err2 := strconv.Atoi(strategyIndicatorRule.Max)
			if err1 != nil || err2 != nil {
				continue
			}
			if min <= renewCount && (renewCount < max || (min > 0 && max <= 0)) {
				score = &users.Score{}
				score.ID = strategyIndicatorRule.ID
				score.Name = strategyIndicatorRule.Name
				score.Score = float64(strategyIndicatorRule.Score)
				return
			}
		}
		return nil, noCalculateScoreError
	},
	"fission_map_count": func(userData *users.Data, strategyIndicatorRules []*strategy.StrategyIndicatorRule) (score *users.Score, err error) {
		fissionMapCount := len(userData.FissionMap)
		for _, strategyIndicatorRule := range strategyIndicatorRules {
			min, err1 := strconv.Atoi(strategyIndicatorRule.Min)
			max, err2 := strconv.Atoi(strategyIndicatorRule.Max)
			if err1 != nil || err2 != nil {
				continue
			}
			if min <= fissionMapCount && (fissionMapCount < max || (min > 0 && max <= 0)) {
				score = &users.Score{}
				score.ID = strategyIndicatorRule.ID
				score.Name = strategyIndicatorRule.Name
				score.Score = float64(strategyIndicatorRule.Score)
				return
			}
		}
		return nil, noCalculateScoreError
	},
	"renew_discount": func(userData *users.Data, strategyIndicatorRules []*strategy.StrategyIndicatorRule) (score *users.Score, err error) {
		var (
			discount  int
			remains   float32
			todayDate = time.TodayDate()
		)
		for _, membership := range userData.Memberships {
			remains += membership.Remains
		}
		for _, couponAlloc := range userData.CouponAllocData.CouponAllocs {
			deadline := couponAlloc.Deadline
			if deadline != time.TimeZeroTime || deadline.Before(todayDate) {
				continue
			}
			couponID := couponAlloc.Coupon.ID
			if couponID == 11 || couponID == 165 {
				discount += int(couponAlloc.Remains)
				continue
			}
			if couponAlloc.Coupon.Type == coupon.CashType && couponAlloc.Coupon.AmountType == coupon.NumberAmountType {
				remains += couponAlloc.Remains
			}
		}
		if remains >= 10 {
			if remains < 20 {
				discount += 500
			} else {
				discount += 1000
			}
		}
		for _, strategyIndicatorRule := range strategyIndicatorRules {
			min, err1 := strconv.Atoi(strategyIndicatorRule.Min)
			max, err2 := strconv.Atoi(strategyIndicatorRule.Max)
			if err1 != nil || err2 != nil {
				continue
			}
			if min <= discount && discount <= max {
				score = &users.Score{}
				score.ID = strategyIndicatorRule.ID
				score.Name = strategyIndicatorRule.Name
				score.Score = float64(strategyIndicatorRule.Score)
				return
			}
		}
		return nil, noCalculateScoreError
	},
	"coupon_remain": func(userData *users.Data, strategyIndicatorRules []*strategy.StrategyIndicatorRule) (score *users.Score, err error) {
		var remains float32
		todayDate := time.TodayDate()
		for _, couponAlloc := range userData.CouponAllocData.CouponAllocs {
			deadline := couponAlloc.Deadline
			if deadline != time.TimeZeroTime || deadline.Before(todayDate) {
				continue
			}
			if couponAlloc.Coupon.Type == coupon.CashType && couponAlloc.Coupon.AmountType == coupon.NumberAmountType {
				remains += couponAlloc.Remains
			}
		}
		var remainsFloat64 float64 = float64(remains)
		for _, strategyIndicatorRule := range strategyIndicatorRules {
			min, err1 := strconv.ParseFloat(strategyIndicatorRule.Min, 64)
			max, err2 := strconv.ParseFloat(strategyIndicatorRule.Max, 64)
			if err1 != nil || err2 != nil {
				continue
			}
			if min <= remainsFloat64 && (remainsFloat64 < max || (min > 0 && max <= 0)) {
				score = &users.Score{}
				score.ID = strategyIndicatorRule.ID
				score.Name = strategyIndicatorRule.Name
				score.Score = float64(strategyIndicatorRule.Score)
				return
			}
		}
		return nil, noCalculateScoreError
	},
	"join_judge_count": func(userData *users.Data, strategyIndicatorRules []*strategy.StrategyIndicatorRule) (score *users.Score, err error) {
		judgeUserCount := len(userData.JudgeUserData)
		for _, strategyIndicatorRule := range strategyIndicatorRules {
			min, err1 := strconv.Atoi(strategyIndicatorRule.Min)
			max, err2 := strconv.Atoi(strategyIndicatorRule.Max)
			if err1 != nil || err2 != nil {
				continue
			}
			if min <= judgeUserCount && (judgeUserCount < max || (min > 0 && max <= 0)) {
				score = &users.Score{}
				score.ID = strategyIndicatorRule.ID
				score.Name = strategyIndicatorRule.Name
				score.Score = float64(strategyIndicatorRule.Score)
				return
			}
		}
		return nil, noCalculateScoreError
	},
	"page_access_count": func(userData *users.Data, strategyIndicatorRules []*strategy.StrategyIndicatorRule) (score *users.Score, err error) {
		pageAccessCount := userData.PageAccessData.PageAccessCount
		for _, strategyIndicatorRule := range strategyIndicatorRules {
			min, err1 := strconv.ParseUint(strategyIndicatorRule.Min, 10, 64)
			max, err2 := strconv.ParseUint(strategyIndicatorRule.Max, 10, 64)
			if err1 != nil || err2 != nil {
				continue
			}
			if min <= pageAccessCount && (pageAccessCount < max || (min > 0 && max <= 0)) {
				score = &users.Score{}
				score.ID = strategyIndicatorRule.ID
				score.Name = strategyIndicatorRule.Name
				score.Score = float64(strategyIndicatorRule.Score)
				return
			}
		}
		return nil, noCalculateScoreError
	},
	"nearly_studio": func(userData *users.Data, strategyIndicatorRules []*strategy.StrategyIndicatorRule) (score *users.Score, err error) {
		signins := userData.Signins
		if len(signins) > 0 {
			var days int

			todayDate := time.TodayDate()
			endDate := todayDate.AddDate(0, 0, 7)

			for index, signin := range signins {
				activityStartDate := time.DateZero(signin.ActivityStartTime)
				if activityStartDate.Before(todayDate) || activityStartDate.After(endDate) {
					continue
				}
				day := int(activityStartDate.Sub(todayDate).Hours() / 24)
				if index == 0 {
					days = day
				} else {
					if day < days {
						days = day
					}
				}
			}
			for _, strategyIndicatorRule := range strategyIndicatorRules {
				min, err1 := strconv.Atoi(strategyIndicatorRule.Min)
				max, err2 := strconv.Atoi(strategyIndicatorRule.Max)
				if err1 != nil || err2 != nil {
					continue
				}
				if min <= days && (days < max || (min > 0 && max <= 0)) {
					score = &users.Score{}
					score.ID = strategyIndicatorRule.ID
					score.Name = strategyIndicatorRule.Name
					score.Score = float64(strategyIndicatorRule.Score)
					return
				}
			}
		}
		return nil, noCalculateScoreError
	},
	"nearly_90_master": func(userData *users.Data, strategyIndicatorRules []*strategy.StrategyIndicatorRule) (score *users.Score, err error) {
		var (
			count     int
			todayDate = time.TodayDate()
			startDate = todayDate.AddDate(0, 0, -90)
			endDate   = todayDate.AddDate(0, 0, 1)
		)

		for _, signin := range userData.Signins {
			activityStartTime := signin.ActivityStartTime
			if signin.CourseLevelID == course_level.Master && activityStartTime.After(startDate) && activityStartTime.Before(endDate) {
				count += 1
			}
		}
		for _, strategyIndicatorRule := range strategyIndicatorRules {
			min, err1 := strconv.Atoi(strategyIndicatorRule.Min)
			max, err2 := strconv.Atoi(strategyIndicatorRule.Max)
			if err1 != nil || err2 != nil {
				continue
			}
			if min <= count && (count < max || (min > 0 && max <= 0)) {
				score = &users.Score{}
				score.ID = strategyIndicatorRule.ID
				score.Name = strategyIndicatorRule.Name
				score.Score = float64(strategyIndicatorRule.Score)
				return
			}
		}
		return nil, noCalculateScoreError
	},
	"nearly_90_xiaoban": func(userData *users.Data, strategyIndicatorRules []*strategy.StrategyIndicatorRule) (score *users.Score, err error) {
		var (
			courseMap = map[uint32]bool{}
			todayDate = time.TodayDate()
			startDate = todayDate.AddDate(0, 0, -90)
			endDate   = todayDate.AddDate(0, 0, 1)
		)

		for _, signin := range userData.Signins {
			activityStartTime := signin.ActivityStartTime
			if signin.CourseLevelID == course_level.XiaoBan && activityStartTime.After(startDate) && activityStartTime.Before(endDate) {
				courseMap[signin.CourseID] = true
			}
		}
		count := len(courseMap)
		for _, strategyIndicatorRule := range strategyIndicatorRules {
			min, err1 := strconv.Atoi(strategyIndicatorRule.Min)
			max, err2 := strconv.Atoi(strategyIndicatorRule.Max)
			if err1 != nil || err2 != nil {
				continue
			}
			if min <= count && (count < max || (min > 0 && max <= 0)) {
				score = &users.Score{}
				score.ID = strategyIndicatorRule.ID
				score.Name = strategyIndicatorRule.Name
				score.Score = float64(strategyIndicatorRule.Score)
				return
			}
		}
		return nil, noCalculateScoreError
	},
	"nearly_90_jixun": func(userData *users.Data, strategyIndicatorRules []*strategy.StrategyIndicatorRule) (score *users.Score, err error) {
		var (
			courseMap = map[uint32]bool{}
			todayDate = time.TodayDate()
			startDate = todayDate.AddDate(0, 0, -90)
			endDate   = todayDate.AddDate(0, 0, 1)
		)

		for _, signin := range userData.Signins {
			activityStartTime := signin.ActivityStartTime
			if signin.CourseLevelID == course_level.JiXun && activityStartTime.After(startDate) && activityStartTime.Before(endDate) {
				courseMap[signin.CourseID] = true
			}
		}
		count := len(courseMap)
		for _, strategyIndicatorRule := range strategyIndicatorRules {
			min, err1 := strconv.Atoi(strategyIndicatorRule.Min)
			max, err2 := strconv.Atoi(strategyIndicatorRule.Max)
			if err1 != nil || err2 != nil {
				continue
			}
			if min <= count && (count < max || (min > 0 && max <= 0)) {
				score = &users.Score{}
				score.ID = strategyIndicatorRule.ID
				score.Name = strategyIndicatorRule.Name
				score.Score = float64(strategyIndicatorRule.Score)
				return
			}
		}
		return nil, noCalculateScoreError
	},
	"show_video_count": func(userData *users.Data, strategyIndicatorRules []*strategy.StrategyIndicatorRule) (score *users.Score, err error) {
		showVideoCount := userData.ShowVideoCount
		for _, strategyIndicatorRule := range strategyIndicatorRules {
			min, err1 := strconv.Atoi(strategyIndicatorRule.Min)
			max, err2 := strconv.Atoi(strategyIndicatorRule.Max)
			if err1 != nil || err2 != nil {
				continue
			}
			if int64(min) <= showVideoCount && (showVideoCount < int64(max) || (min > 0 && max <= 0)) {
				score = &users.Score{}
				score.ID = strategyIndicatorRule.ID
				score.Name = strategyIndicatorRule.Name
				score.Score = float64(strategyIndicatorRule.Score)
				return
			}
		}
		return nil, noCalculateScoreError
	},
	"all_signin_spend": func(userData *users.Data, strategyIndicatorRules []*strategy.StrategyIndicatorRule) (score *users.Score, err error) {
		allSigninSpend := userData.AllSigninSpend
		for _, strategyIndicatorRule := range strategyIndicatorRules {
			min, err1 := strconv.ParseFloat(strategyIndicatorRule.Min, 64)
			max, err2 := strconv.ParseFloat(strategyIndicatorRule.Max, 64)
			if err1 != nil || err2 != nil {
				continue
			}
			if min <= allSigninSpend && (allSigninSpend < max || (min > 0 && max <= 0)) {
				score = &users.Score{}
				score.ID = strategyIndicatorRule.ID
				score.Name = strategyIndicatorRule.Name
				score.Score = float64(strategyIndicatorRule.Score)
				return
			}
		}
		return nil, noCalculateScoreError
	},
	"total_membership_amount": func(userData *users.Data, strategyIndicatorRules []*strategy.StrategyIndicatorRule) (score *users.Score, err error) {
		var amount float64
		for _, order := range userData.Orders {
			orderStatus := order.Status
			if !(orderStatus == orderRepo.STATUS_WAIT_SEND || orderStatus == orderRepo.STATUS_WAIT_RECEIVE || orderStatus == orderRepo.STATUS_COMPLETE) {
				continue
			}
			for _, orderItem := range order.OrderItems {
				if orderItem.Type != 1 {
					continue
				}
				amount += float64(orderItem.Price)
			}
		}
		for _, strategyIndicatorRule := range strategyIndicatorRules {
			min, err1 := strconv.ParseFloat(strategyIndicatorRule.Min, 64)
			max, err2 := strconv.ParseFloat(strategyIndicatorRule.Max, 64)
			if err1 != nil || err2 != nil {
				continue
			}
			if min <= amount && (amount < max || (min > 0 && max <= 0)) {
				score = &users.Score{}
				score.ID = strategyIndicatorRule.ID
				score.Name = strategyIndicatorRule.Name
				score.Score = float64(strategyIndicatorRule.Score)
				return
			}
		}
		return nil, noCalculateScoreError
	},
	// --------------------------------------------------------------
	"clues_transfer_time": func(userData *users.Data, strategyIndicatorRules []*strategy.StrategyIndicatorRule) (score *users.Score, err error) {
		transferDate := time.DateZero(userData.UserBeforeMember.TransferTime)
		todayDate := time.TodayDate()
		day := int(todayDate.Sub(transferDate).Hours() / 24)
		for _, strategyIndicatorRule := range strategyIndicatorRules {
			min, err1 := strconv.Atoi(strategyIndicatorRule.Min)
			max, err2 := strconv.Atoi(strategyIndicatorRule.Max)
			if err1 != nil || err2 != nil {
				continue
			}
			if min <= day && (day < max || (min > 0 && max <= 0)) {
				score = &users.Score{}
				score.ID = strategyIndicatorRule.ID
				score.Name = strategyIndicatorRule.Name
				score.Score = float64(strategyIndicatorRule.Score)
				return
			}
		}
		return nil, noCalculateScoreError
	},
	"clues_source": func(userData *users.Data, strategyIndicatorRules []*strategy.StrategyIndicatorRule) (score *users.Score, err error) {
		sourceType := source.SourceType[userData.UserBeforeMember.SourceID]
		for _, strategyIndicatorRule := range strategyIndicatorRules {
			score = &users.Score{}
			if sourceType == strategyIndicatorRule.Min {
				score.ID = strategyIndicatorRule.ID
				score.Name = strategyIndicatorRule.Name
				score.Score = float64(strategyIndicatorRule.Score)
				return
			}
		}
		return nil, noCalculateScoreError
	},
	"last_page_access": func(userData *users.Data, strategyIndicatorRules []*strategy.StrategyIndicatorRule) (score *users.Score, err error) {
		lastPageAccessTime := userData.PageAccessData.LastPageAccessTime
		if lastPageAccessTime != time.TimeZeroTime {
			todayDate := time.TodayDate()
			day := int(todayDate.Sub(time.DateZero(lastPageAccessTime)).Hours() / 24)
			for _, strategyIndicatorRule := range strategyIndicatorRules {
				min, err1 := strconv.Atoi(strategyIndicatorRule.Min)
				max, err2 := strconv.Atoi(strategyIndicatorRule.Max)
				if err1 != nil || err2 != nil {
					continue
				}
				if min <= day && (day < max || (min > 0 && max <= 0)) {
					score = &users.Score{}
					score.ID = strategyIndicatorRule.ID
					score.Name = strategyIndicatorRule.Name
					score.Score = float64(strategyIndicatorRule.Score)
					return
				}
			}
		}
		return nil, noCalculateScoreError
	},
	// --------------------------------------------------------------
	"coupon_days": func(userData *users.Data, strategyIndicatorRules []*strategy.StrategyIndicatorRule) (score *users.Score, err error) {
		var couponAllocCreateTime time.Time
		for _, couponAlloc := range userData.CouponAllocData.CouponAllocs {
			if couponAlloc.Coupon.IsNewUser {
				couponAllocCreateTime = couponAlloc.CreateTime
			}
		}
		if couponAllocCreateTime != time.TimeZeroTime {
			todayDate := time.TodayDate()
			day := int(todayDate.Sub(time.DateZero(couponAllocCreateTime)).Hours() / 24)
			for _, strategyIndicatorRule := range strategyIndicatorRules {
				min, err1 := strconv.Atoi(strategyIndicatorRule.Min)
				max, err2 := strconv.Atoi(strategyIndicatorRule.Max)
				if err1 != nil || err2 != nil {
					continue
				}
				if min <= day && (day < max || (min > 0 && max <= 0)) {
					score = &users.Score{}
					score.ID = strategyIndicatorRule.ID
					score.Name = strategyIndicatorRule.Name
					score.Score = float64(strategyIndicatorRule.Score)
					return
				}
			}
		}
		return nil, noCalculateScoreError
	},
	"access_course_detail": func(userData *users.Data, strategyIndicatorRules []*strategy.StrategyIndicatorRule) (score *users.Score, err error) {
		currentStudioAccessActivityCount := userData.PageAccessData.CurrentStudioAccessActivityCount
		for _, strategyIndicatorRule := range strategyIndicatorRules {
			min, err1 := strconv.ParseUint(strategyIndicatorRule.Min, 10, 64)
			max, err2 := strconv.ParseUint(strategyIndicatorRule.Max, 10, 64)
			if err1 != nil || err2 != nil {
				continue
			}
			if min <= currentStudioAccessActivityCount && (currentStudioAccessActivityCount < max || (min > 0 && max <= 0)) {
				score = &users.Score{}
				score.ID = strategyIndicatorRule.ID
				score.Name = strategyIndicatorRule.Name
				score.Score = float64(strategyIndicatorRule.Score)
				return
			}
		}
		return nil, noCalculateScoreError
	},
	"geolocation": func(userData *users.Data, strategyIndicatorRules []*strategy.StrategyIndicatorRule) (score *users.Score, err error) {
		accessLocation := userData.PageEventData.AccessLocation
		longitude, err1 := strconv.ParseFloat(accessLocation.Longitude, 64)
		latitude, err2 := strconv.ParseFloat(accessLocation.Latitude, 64)
		if err1 == nil && err2 == nil {
			distance := utils.Distance(longitude, latitude, utils.TianAnMen.Longitude, utils.TianAnMen.Latitude)
			var geolocation string
			if distance <= utils.BeiJindRadius {
				geolocation = "in"
			} else {
				geolocation = "notin"
			}
			for _, strategyIndicatorRule := range strategyIndicatorRules {
				if geolocation == strategyIndicatorRule.Min {
					score = &users.Score{}
					score.ID = strategyIndicatorRule.ID
					score.Name = strategyIndicatorRule.Name
					score.Score = float64(strategyIndicatorRule.Score)
					return
				}
			}
		}
		return nil, noCalculateScoreError
	},
	"follow_time": func(userData *users.Data, strategyIndicatorRules []*strategy.StrategyIndicatorRule) (score *users.Score, err error) {
		var couponAllocCreateTime time.Time

		for _, couponAlloc := range userData.CouponAllocData.CouponAllocs {
			if couponAlloc.Coupon.IsNewUser {
				couponAllocCreateTime = couponAlloc.CreateTime
			}
		}
		if couponAllocCreateTime == time.TimeZeroTime {
			return nil, noCalculateScoreError
		}
		couponAllocCreateTime = time.DateZero(couponAllocCreateTime)

		var day int
		var haveFollow bool
		for _, follow := range userData.UserBeforeMember.Follows {
			followCreateDate := time.DateZero(follow.CreateTime)
			if followCreateDate.After(couponAllocCreateTime) || followCreateDate.Equal(couponAllocCreateTime) {
				day = int(followCreateDate.Sub(couponAllocCreateTime).Hours() / 24)
				haveFollow = true
				break
			}
		}
		if !haveFollow {
			return nil, noCalculateScoreError
		}

		for _, strategyIndicatorRule := range strategyIndicatorRules {
			min, err1 := strconv.Atoi(strategyIndicatorRule.Min)
			max, err2 := strconv.Atoi(strategyIndicatorRule.Max)
			if err1 != nil || err2 != nil {
				continue
			}
			if min <= day && (day < max || (min > 0 && max <= 0)) {
				score = &users.Score{}
				score.ID = strategyIndicatorRule.ID
				score.Name = strategyIndicatorRule.Name
				score.Score = float64(strategyIndicatorRule.Score)
				return
			}
		}
		return nil, noCalculateScoreError
	},
}

/*
计算某个指标的分数
*/
func StrategyIndicatorCalculate(strategyIndicator *strategy.StrategyIndicator, userData *users.Data) (score *users.Score, err error) {
	calculateFunc := strategyIndicatorCalculateFunc[strategyIndicator.Key]
	if calculateFunc == nil {
		return nil, errors.New(fmt.Sprintf("无【%s】指标计算逻辑", strategyIndicator.Key))
	}
	score, err = calculateFunc(userData, strategyIndicator.StrategyIndicatorRules)
	if err != nil {
		return
	}
	score.Weight = strategyIndicator.Weight
	score.Name = strategyIndicator.Name + ": " + score.Name
	score.Score = score.Score * float64(strategyIndicator.Weight) / 100
	return
}
