package strategy_indicator

import (
	"fmt"
	"signin-go/global/time"
	"signin-go/internal/errors"
	"signin-go/repository/coupon"
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
