package todo

import (
	"time"
)

// PaymentAccount 收款平台，下的收款账户
type PaymentAccount struct {
	ID                uint32    `gorm:"primaryKey;column:id" json:"id"`
	PaymentPlatformID uint32    `gorm:"column:payment_platform_id" json:"payment_platform_id"` // 对应的收款平台id
	PaymentBankID     uint32    `gorm:"column:payment_bank_id" json:"payment_bank_id"`         // 对应银行账号
	Name              string    `gorm:"column:name" json:"name"`                               // 名称，自有微信支付、招行平台、自有支付宝平台、现金
	Status            uint32    `gorm:"column:status" json:"status"`                           // 状态，1有效，2暂停使用
	MerID             string    `gorm:"column:merId" json:"mer_id"`                            // 商户号
	Mark              string    `gorm:"column:mark" json:"mark"`                               // 备注
	IsDel             bool      `gorm:"column:is_del" json:"is_del"`
	CreateTime        time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime        time.Time `gorm:"column:update_time" json:"update_time"`
	StudioNo          string    `gorm:"column:studio_no" json:"studio_no"` // 门店编号
	AppID             string    `gorm:"column:app_id" json:"app_id"`
	AppSecret         string    `gorm:"column:app_secret" json:"app_secret"`
	CashierNo         string    `gorm:"column:cashier_no" json:"cashier_no"` // 收银员编号
}

// TableName get sql table name.获取数据库表名
func (m *PaymentAccount) TableName() string {
	return "payment_account"
}
