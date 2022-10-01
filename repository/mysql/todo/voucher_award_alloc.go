package todo

import (
	"time"
)

// VoucherAwardAlloc 兑换的奖品实例对应
type VoucherAwardAlloc struct {
	ID         uint32    `gorm:"primaryKey;column:id" json:"id"`
	VoucherID  uint32    `gorm:"column:voucher_id" json:"voucher_id"`   // 兑换码id
	Type       uint32    `gorm:"column:type" json:"type"`               // 奖品类型，1会员卡，2课程，3优惠券, 4sku 商品
	ItemID     uint32    `gorm:"column:item_id" json:"item_id"`         // 奖品id(券模板id、课程id、sku_id 等)
	InstanceID uint32    `gorm:"column:instance_id" json:"instance_id"` // 奖品的实例id, 例如约课的signin_id 等
	Status     uint32    `gorm:"column:status" json:"status"`           // 备用，状态，1正常，2已退
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"`
	IsDel      bool      `gorm:"column:is_del" json:"is_del"`
}

// TableName get sql table name.获取数据库表名
func (m *VoucherAwardAlloc) TableName() string {
	return "voucher_award_alloc"
}
