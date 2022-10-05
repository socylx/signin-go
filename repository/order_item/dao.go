package order_item

import (
	"signin-go/global/mysql"
	"signin-go/internal/core"
)

func GetOrderIemDatas(ctx core.StdContext, orderID uint32) (data []*OrderItemData, err error) {
	db := mysql.DB.WithContext(ctx)
	err = db.Table("order_item").
		Select("order_item.id,order_item.type,order_item.price").
		Where("order_item.is_del = 0 AND order_item.order_id = ?", orderID).
		Find(&data).Error
	return
}
