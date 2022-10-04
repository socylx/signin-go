package strategy

import (
	"signin-go/global/mysql"
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
	err = db.Table(tableName()).
		Where("strategy.is_del = 0 AND strategy.id = ?", filter.ID).
		Update("strategy.is_del", 1).Error

	return
}

func Detail(ctx core.StdContext, strategyID uint32) (strategy *Strategy, err error) {
	db := mysql.DB.WithContext(ctx)

	strategy = &Strategy{}
	err = db.Table(tableName()).
		Where("strategy.is_del = 0 AND strategy.id = ?", strategyID).
		First(strategy).Error
	return
}
