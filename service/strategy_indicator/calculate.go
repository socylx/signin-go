package strategy_indicator

import (
	"fmt"
	"signin-go/internal/errors"
	"signin-go/repository/strategy"
	"signin-go/repository/users"
)

type CalculateFunc func(userData *users.Data, strategyIndicatorRules []*strategy.StrategyIndicatorRule) (score *users.Score, err error)

var strategyIndicatorCalculateFunc = map[string]CalculateFunc{}

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
