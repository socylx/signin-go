package todo

import (
	"time"
)

// ActivityGroupItem [...]
type ActivityGroupItem struct {
	ID              uint32    `gorm:"primaryKey;column:id" json:"id"`
	ActivityGroupID uint32    `gorm:"column:activity_group_id" json:"activity_group_id"`
	ItemID          uint32    `gorm:"column:item_id" json:"item_id"`
	ItemType        uint32    `gorm:"column:item_type" json:"item_type"`           // 1-课程,2-录播课
	Title           string    `gorm:"column:title" json:"title"`                   // 标题
	Desc            string    `gorm:"column:desc" json:"desc"`                     // 课程描述
	Index           uint32    `gorm:"column:index" json:"index"`                   // 排序值
	Notice          string    `gorm:"column:notice" json:"notice"`                 // 公告
	PriceAmount     float32   `gorm:"column:price_amount" json:"price_amount"`     // 现金价格
	PriceDiscount   float32   `gorm:"column:price_discount" json:"price_discount"` // 如果已经购买过这节课，优惠价格
	CreateTime      time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime      time.Time `gorm:"column:update_time" json:"update_time"`
	IsDel           bool      `gorm:"column:is_del" json:"is_del"`
}

// TableName get sql table name.获取数据库表名
func (m *ActivityGroupItem) TableName() string {
	return "activity_group_item"
}
