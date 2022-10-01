package todo

import (
	"time"
)

// OpinionSuggestion [...]
type OpinionSuggestion struct {
	ID           uint32    `gorm:"primaryKey;column:id" json:"id"`
	UserID       uint32    `gorm:"column:user_id" json:"user_id"`             // 反馈人user_id
	Content      string    `gorm:"column:content" json:"content"`             // 内容
	Contact      string    `gorm:"column:contact" json:"contact"`             // 电话/微信/邮箱等
	AllowContact bool      `gorm:"column:allow_contact" json:"allow_contact"` // 是否允许工作人员联系
	Type         uint32    `gorm:"column:type" json:"type"`                   // 1、反馈、意见、建议, 2、投诉
	Status       uint32    `gorm:"column:status" json:"status"`               // 1、未处理, 2、已处理
	ForID        uint32    `gorm:"column:for_id" json:"for_id"`               // 反馈对应的实例id，例如课程id
	Score        uint32    `gorm:"column:score" json:"score"`                 // 打分
	Remark       string    `gorm:"column:remark" json:"remark"`               // 工作人员处理备注
	OptUserID    uint32    `gorm:"column:opt_user_id" json:"opt_user_id"`     // 处理人员的user_id
	CreateTime   time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime   time.Time `gorm:"column:update_time" json:"update_time"`
	IsDel        bool      `gorm:"column:is_del" json:"is_del"`
}

// TableName get sql table name.获取数据库表名
func (m *OpinionSuggestion) TableName() string {
	return "opinion_suggestion"
}
