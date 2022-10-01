package todo

import (
	"time"
)

// StudioPaymentAccount 门店对应的支付账户，n:n
type StudioPaymentAccount struct {
	ID               uint32    `gorm:"primaryKey;column:id" json:"id"`
	StudioID         uint32    `gorm:"column:studio_id" json:"studio_id"`                   // 门店id
	PaymentAccountID uint32    `gorm:"column:payment_account_id" json:"payment_account_id"` // 支付账户id
	IsDel            bool      `gorm:"column:is_del" json:"is_del"`
	CreateTime       time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime       time.Time `gorm:"column:update_time" json:"update_time"`
}

// TableName get sql table name.获取数据库表名
func (m *StudioPaymentAccount) TableName() string {
	return "studio_payment_account"
}
