package todo

import (
	"time"
)

// ShopGroupAlloc [...]
type ShopGroupAlloc struct {
	ID           uint32    `gorm:"primaryKey;column:id" json:"id"`
	ShopGroupID  uint32    `gorm:"column:shop_group_id" json:"shop_group_id"`   // 团购活动id
	MasterUserID uint32    `gorm:"column:master_user_id" json:"master_user_id"` // 拼团发起人id
	Status       uint32    `gorm:"column:status" json:"status"`                 // 本次拼团状态，1初始化，2超时，3未成团，4成功
	EndTime      time.Time `gorm:"column:end_time" json:"end_time"`             // 本次拼团截止时间
	IsDel        bool      `gorm:"column:is_del" json:"is_del"`
	UpdateTime   time.Time `gorm:"column:update_time" json:"update_time"`
	CreateTime   time.Time `gorm:"column:create_time" json:"create_time"`
}

// TableName get sql table name.获取数据库表名
func (m *ShopGroupAlloc) TableName() string {
	return "shop_group_alloc"
}
