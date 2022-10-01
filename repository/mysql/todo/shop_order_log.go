package todo

import (
	"time"
)

// ShopOrderLog 订单日志
type ShopOrderLog struct {
	ID         uint32    `gorm:"primaryKey;column:id" json:"id"`
	OrderID    uint32    `gorm:"column:order_id" json:"order_id"`       // 订单id
	Content    string    `gorm:"column:content" json:"content"`         // 日志内容
	SysContent string    `gorm:"column:sys_content" json:"sys_content"` // 系统自动生成的日志
	OptUserID  uint32    `gorm:"column:opt_user_id" json:"opt_user_id"`
	IsDel      bool      `gorm:"column:is_del" json:"is_del"`
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"`
}

// TableName get sql table name.获取数据库表名
func (m *ShopOrderLog) TableName() string {
	return "shop_order_log"
}
