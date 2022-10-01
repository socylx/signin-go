package todo

import (
	"time"
)

// VoucherTemplate 兑换码模板，定义一些规则
type VoucherTemplate struct {
	ID         uint32    `gorm:"primaryKey;column:id" json:"id"`
	Type       uint32    `gorm:"column:type" json:"type"`     // 活动类型标记，1这街啤酒
	Limit      uint32    `gorm:"column:limit" json:"limit"`   // 每个人限制领取次数
	Mark       string    `gorm:"column:mark" json:"mark"`     // 备注
	Notice     string    `gorm:"column:notice" json:"notice"` // 活动通告，说明
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"`
	IsDel      bool      `gorm:"column:is_del" json:"is_del"`
	BgURL      string    `gorm:"column:bg_url" json:"bg_url"`           // 活动的头图
	ImgURL1    string    `gorm:"column:img_url_1" json:"img_url_1"`     // 通用图片1
	SourceID   uint32    `gorm:"column:source_id" json:"source_id"`     // 兑换码设置来源
	CodeLength uint32    `gorm:"column:code_length" json:"code_length"` // 生成的兑换码长度
}

// TableName get sql table name.获取数据库表名
func (m *VoucherTemplate) TableName() string {
	return "voucher_template"
}
