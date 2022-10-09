package order

import (
	"gsteps-go/global/mysql"
	"gsteps-go/internal/core"
)

func GetOrderDatas(ctx core.StdContext, userID uint32) (data []*OrderData, err error) {
	db := mysql.DB.WithContext(ctx)
	err = db.Table("order").
		Select("order.id, order.status").
		Where("order.is_del = 0 AND order.user_id = ?", userID).
		Find(&data).Error
	return
}
