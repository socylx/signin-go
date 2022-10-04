package strategy

import (
	"signin-go/global/mysql"
	"signin-go/global/time"
	"signin-go/internal/core"
)

type DeleteFileter struct {
	ID uint32
}

func Delete(ctx core.StdContext, filter *DeleteFileter) (err error) {
	if filter.ID <= 0 {
		return
	}

	db := mysql.DB.WithContext(ctx)
	err = db.Table("strategy").
		Where("strategy.is_del = 0 AND strategy.id = ?", filter.ID).
		Update("strategy.is_del", 1).Error

	return
}

func Detail(ctx core.StdContext, strategyID uint32) (strategy *Strategy, err error) {
	db := mysql.DB.WithContext(ctx)

	strategy = &Strategy{}
	err = db.Table("strategy").
		Where("strategy.is_del = 0 AND strategy.id = ?", strategyID).
		First(strategy).Error
	return
}

type ListFilter struct {
	IncludeIds   []uint32
	CreateTimeGE time.Time
	CreateTimeLT time.Time
	Keyword      string
	Status       int
	CreateUserID uint32
	Type         uint32
	Page         int
	Size         int
}

type listResult struct {
	Data  []*Strategy `json:"data"`
	Count int64       `json:"count"`
}

func List(ctx core.StdContext, filter *ListFilter) (result *listResult, err error) {
	db := mysql.DB.WithContext(ctx)

	query := db.Table("strategy").Where("strategy.is_del = 0")
	if len(filter.IncludeIds) > 0 {
		query = query.Where("strategy.id IN ?", filter.IncludeIds)
	}
	if filter.CreateTimeGE != time.TimeZeroTime {
		query = query.Where("strategy.create_time >= ?", filter.CreateTimeGE)
	}
	if filter.CreateTimeLT != time.TimeZeroTime {
		query = query.Where("strategy.create_time < ?", filter.CreateTimeLT.Add(24*time.Hour))
	}
	if filter.Keyword != "" {
		likeValue := "%" + filter.Keyword + "%"
		query = query.Or("(strategy.name LIKE ? OR strategy.desc LIKE ?)", likeValue, likeValue)
	}
	if filter.Status > 0 {
		query = query.Where("strategy.status = ?", filter.Status)
	}
	if filter.CreateUserID > 0 {
		query = query.Where("strategy.create_user_id = ?", filter.CreateUserID)
	}
	if filter.Type > 0 {
		query = query.Where("strategy.type = ?", filter.Type)
	}
	query.Count(&result.Count)
	err = query.Order("strategy.id DESC").
		Limit(filter.Size).
		Offset(filter.Page*filter.Size - filter.Size).
		Find(&result.Data).Error
	return
}

func GetStrategyIndicatorDatas(ctx core.StdContext, strategyIndicatorIDs []uint32) (data []*IndicatorData, err error) {
	db := mysql.DB.WithContext(ctx)
	db.Table("strategy_indicator").
		Select("strategy_indicator.*").
		Where("strategy_indicator.is_del = 0 AND strategy_indicator.id IN ?", strategyIndicatorIDs).
		Find(&data)
	return
}

func GetStrategyIndicatorRuleDatas(ctx core.StdContext, strategyIndicatorRuleIDs []uint32) (data []*IndicatorRuleData, err error) {
	db := mysql.DB.WithContext(ctx)

	err = db.Table("strategy_indicator_rule").
		Select("strategy_indicator_rule.*").
		Where("strategy_indicator_rule.is_del = 0 AND strategy_indicator_rule.id IN ?", strategyIndicatorRuleIDs).
		Find(&data).Error
	return
}
