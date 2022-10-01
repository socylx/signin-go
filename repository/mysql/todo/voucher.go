package todo

import (
	"time"
)

// Voucher 兑换码实例
type Voucher struct {
	ID                uint32    `gorm:"primaryKey;column:id" json:"id"`
	VoucherTemplateID uint32    `gorm:"column:voucher_template_id" json:"voucher_template_id"` // 所属活动id
	Code              string    `gorm:"column:code" json:"code"`                               // 兑换码，都存大写字母
	Status            uint32    `gorm:"column:status" json:"status"`                           // 状态，1初始化，2已兑换，3失效
	UserID            uint32    `gorm:"column:user_id" json:"user_id"`                         // 用户id
	CreateTime        time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime        time.Time `gorm:"column:update_time" json:"update_time"`
	IsDel             bool      `gorm:"column:is_del" json:"is_del"`
}

// TableName get sql table name.获取数据库表名
func (m *Voucher) TableName() string {
	return "voucher"
}
