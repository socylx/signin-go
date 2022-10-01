package todo

import (
	"time"
)

// PaymentBank 公司银行账户信息
type PaymentBank struct {
	ID                uint32    `gorm:"primaryKey;column:id" json:"id"`
	Name              string    `gorm:"column:name" json:"name"`                               // 用户账户名称
	Key               string    `gorm:"column:key" json:"key"`                                 // 银行的key, cmb招行，icbc工商银行
	Account           string    `gorm:"column:account" json:"account"`                         // 银行账号
	BranchAccount     string    `gorm:"column:branch_account" json:"branch_account"`           // 分行号
	Mark              string    `gorm:"column:mark" json:"mark"`                               // 备注
	BranchAccountName string    `gorm:"column:branch_account_name" json:"branch_account_name"` // 分行、支行名称
	Status            uint32    `gorm:"column:status" json:"status"`                           // 状态，1有效，2暂停使用
	IsDel             bool      `gorm:"column:is_del" json:"is_del"`
	CreateTime        time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime        time.Time `gorm:"column:update_time" json:"update_time"`
	Data1             string    `gorm:"column:data_1" json:"data_1"`     // 备用字段1
	Data2             string    `gorm:"column:data_2" json:"data_2"`     // 备用字段3
	Data3             string    `gorm:"column:data_3" json:"data_3"`     // 备用字段3
	DataKey           string    `gorm:"column:data_key" json:"data_key"` // 备用字段
	DataPri           string    `gorm:"column:data_pri" json:"data_pri"` // 备用字段
	DataPub           string    `gorm:"column:data_pub" json:"data_pub"` // 备用字段
	DataUID           string    `gorm:"column:data_uid" json:"data_uid"` // 备用字段
}

// TableName get sql table name.获取数据库表名
func (m *PaymentBank) TableName() string {
	return "payment_bank"
}
