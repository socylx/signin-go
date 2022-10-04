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
		Where("strategy.is_del = 0 AND strategy.ID = ?", filter.ID).
		Update("strategy.is_del", 1).Error

	return
}
