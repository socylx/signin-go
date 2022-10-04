package strategy

import (
	"crypto/md5"
	"fmt"
	"signin-go/global/time"
	"signin-go/internal/core"
	"signin-go/repository/strategy"
)

const strategyKeyBaseString = "30ec877eaf21e960b504398cc7f48efc"

func GenerateStrategyKey() string {
	data := []byte(fmt.Sprintf("%s %s", strategyKeyBaseString, time.CSTLayoutString(time.Now(), time.CSTLayout)))
	sumStr := fmt.Sprintf("%x", md5.Sum(data))
	return sumStr
}

// func MatchUserStrategyType(ctx core.StdContext, userData map[string]interface{}) (strategyType uint32) {
// 	memberships := userData["memberships"].([]map[string]interface{})
// 	if len(memberships) > 0 {
// 		strategyType = strategy.StrategyType.Xuka
// 	} else {
// 		couponAllocs := userData["coupon_alloc_data"].(*user.CouponAllocData).CouponAllocs

// 		var couponAllocID uint32
// 		for _, couponAlloc := range couponAllocs {
// 			if couponAlloc["is_new_user"].(bool) {
// 				couponAllocID = couponAlloc["id"].(uint32)
// 			}
// 		}
// 		if couponAllocID > 0 {
// 			signin, _ := h.signinService.DetailByCouponAllocID(c, couponAllocID)

// 			if signin.ID <= 0 {
// 				strategyType = strategy.StrategyType.LaXinSubscribe
// 			} else {
// 				if signin.ActivityStartTime.Before(time.Now()) {
// 					strategyType = strategy.StrategyType.LaXinTry
// 				} else {
// 					strategyType = strategy.StrategyType.LaXinMembership
// 				}
// 			}
// 		} else {
// 			strategyType = strategy.StrategyType.LaXinCoupon
// 		}
// 	}
// 	return
// }

func Data(ctx core.StdContext, strategyID uint32) (data *strategy.StrategyDocument, err error) {
	s, err := strategy.Detail(ctx, strategyID)
	if err != nil {
		return
	}

	var strategyIndicators []*strategy.StrategyIndicator

	strategyIndicatorDatas, _ := IndicatorDataList(ctx, strategyID)
	for _, strategyIndicatorData := range strategyIndicatorDatas {
		var strategyIndicatorRules []*strategy.StrategyIndicatorRule
		for _, strategyIndicatorRuleData := range strategyIndicatorData.Rules {
			strategyIndicatorRules = append(
				strategyIndicatorRules,
				&strategy.StrategyIndicatorRule{
					ID:         strategyIndicatorRuleData.ID,
					Name:       strategyIndicatorRuleData.Name,
					Min:        strategyIndicatorRuleData.Min,
					Max:        strategyIndicatorRuleData.Max,
					CreateTime: strategyIndicatorRuleData.CreateTime,
					UpdateTime: strategyIndicatorRuleData.UpdateTime,
					IsDel:      strategyIndicatorRuleData.IsDel,
					Score:      strategyIndicatorRuleData.Score,
				},
			)
		}

		strategyIndicators = append(
			strategyIndicators,
			&strategy.StrategyIndicator{
				ID:                     strategyIndicatorData.ID,
				Key:                    strategyIndicatorData.Key,
				Name:                   strategyIndicatorData.Name,
				CreateTime:             strategyIndicatorData.CreateTime,
				UpdateTime:             strategyIndicatorData.UpdateTime,
				IsDel:                  strategyIndicatorData.IsDel,
				Weight:                 strategyIndicatorData.Weight,
				StrategyIndicatorRules: strategyIndicatorRules,
			},
		)
	}
	data = &strategy.StrategyDocument{
		ID:                 s.ID,
		Name:               s.Name,
		Desc:               s.Desc,
		Type:               s.Type,
		Status:             s.Status,
		StartTime:          s.StartTime,
		Key:                s.Key,
		CreateUserID:       s.CreateUserID,
		CreateTime:         s.CreateTime,
		UpdateTime:         s.UpdateTime,
		IsDel:              s.IsDel,
		StrategyIndicators: strategyIndicators,
	}
	return
}
