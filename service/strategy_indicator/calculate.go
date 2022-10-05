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
				remains += couponAlloc.Remain
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
