package todo

import (
	"time"
)

// JudgeRule [...]
type JudgeRule struct {
	ID         uint32    `gorm:"primaryKey;column:id" json:"id"`
	JudgeID    uint32    `gorm:"column:judge_id" json:"judge_id"`
	Type       uint32    `gorm:"column:type" json:"type"` // 类型，1有文案，2不含图片，3含图片，4用户类型，5海报key，6成为会员时间大于等于，7成为会员时间小于等于
	Value      string    `gorm:"column:value" json:"value"`
	GroupNo    uint32    `gorm:"column:group_no" json:"group_no"` // 规则所属组，相同组直接【或】关系，不同组之间【且】关系
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"`
	IsDel      bool      `gorm:"column:is_del" json:"is_del"`
	ErrMsg     string    `gorm:"column:err_msg" json:"err_msg"` // 不满足时，展示给用户的文案
}

// TableName get sql table name.获取数据库表名
func (m *JudgeRule) TableName() string {
	return "judge_rule"
}
