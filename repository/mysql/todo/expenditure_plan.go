package todo

import (
	"time"
)

// ExpenditurePlan 支出计划表
type ExpenditurePlan struct {
	ID                         uint32    `gorm:"primaryKey;column:id" json:"id"`
	UserID                     uint32    `gorm:"column:user_id" json:"user_id"`                                               // 计划对应人
	Deadline                   time.Time `gorm:"column:deadline" json:"deadline"`                                             // 截止时间
	StudioID                   uint32    `gorm:"column:studio_id" json:"studio_id"`                                           // 门店
	PaymentBankID              uint32    `gorm:"column:payment_bank_id" json:"payment_bank_id"`                               // 银行账户id
	Expenditure                float32   `gorm:"column:expenditure" json:"expenditure"`                                       // 支出金额
	Status                     uint32    `gorm:"column:status" json:"status"`                                                 // 1未支付，2已支付，3废弃
	ExpenditureTypeID          uint32    `gorm:"column:expenditure_type_id" json:"expenditure_type_id"`                       // 用途id
	Mark                       string    `gorm:"column:mark" json:"mark"`                                                     // 备注
	CreateTime                 time.Time `gorm:"column:create_time" json:"create_time"`                                       // 创建时间
	UpdateTime                 time.Time `gorm:"column:update_time" json:"update_time"`                                       // 更新时间
	OptUserID                  uint32    `gorm:"column:opt_user_id" json:"opt_user_id"`                                       // 创建人
	RepeatType                 uint32    `gorm:"column:repeat_type" json:"repeat_type"`                                       // 循环支持类型，1不循环，2日循环，3周循环，4月循环
	RepeatValue                uint32    `gorm:"column:repeat_value" json:"repeat_value"`                                     // 循环值，例如 repeat_type=3，repeat_value=2 则表示2周一循环
	RepeatStatus               uint32    `gorm:"column:repeat_status" json:"repeat_status"`                                   // 1正常，2已派生出下一个
	PreRepeatExpenditurePlanID uint32    `gorm:"column:pre_repeat_expenditure_plan_id" json:"pre_repeat_expenditure_plan_id"` // 本次支出计划，是由哪个计划生成的
	IsDel                      bool      `gorm:"column:is_del" json:"is_del"`
}

// TableName get sql table name.获取数据库表名
func (m *ExpenditurePlan) TableName() string {
	return "expenditure_plan"
}
