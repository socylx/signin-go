package todo

import (
	"time"
)

// UserProfile [...]
type UserProfile struct {
	ID               uint32    `gorm:"primaryKey;column:id" json:"id"`
	UserID           uint32    `gorm:"column:user_id" json:"user_id"`
	Sex              uint32    `gorm:"column:sex" json:"sex"`
	Phone            string    `gorm:"column:phone" json:"phone"`
	PayeeName        string    `gorm:"column:payee_name" json:"payee_name"`                 // 真实姓名，用于工资发放
	PayeeAccount     string    `gorm:"column:payee_account" json:"payee_account"`           // 收款账号
	AccountBelongsTo string    `gorm:"column:account_belongs_to" json:"account_belongs_to"` // 账号所属，支付宝/微信/开户行
	CreateTime       time.Time `gorm:"column:create_time" json:"create_time"`               // 创建时间
	UpdateTime       time.Time `gorm:"column:update_time" json:"update_time"`               // 更新时间
	IsDel            bool      `gorm:"column:is_del" json:"is_del"`                         // 1代表删除
}

// TableName get sql table name.获取数据库表名
func (m *UserProfile) TableName() string {
	return "user_profile"
}
