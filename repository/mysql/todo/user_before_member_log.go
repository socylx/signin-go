package todo

import (
	"time"
)

// UserBeforeMemberLog 售前用户表，操作日志，修改了哪些字段记录
type UserBeforeMemberLog struct {
	ID           uint32    `gorm:"primaryKey;column:id" json:"id"`
	UserBeforeID uint32    `gorm:"column:user_before_id" json:"user_before_id"` // 售前客户id
	OptUserID    int       `gorm:"column:opt_user_id" json:"opt_user_id"`       // 操作人id
	SysMark      string    `gorm:"column:sys_mark" json:"sys_mark"`             // 系统记录的日志内容
	Mark         string    `gorm:"column:mark" json:"mark"`                     // 备注
	IsDel        bool      `gorm:"column:is_del" json:"is_del"`
	UpdateTime   time.Time `gorm:"column:update_time" json:"update_time"`
	CreateTime   time.Time `gorm:"column:create_time" json:"create_time"`
}

// TableName get sql table name.获取数据库表名
func (m *UserBeforeMemberLog) TableName() string {
	return "user_before_member_log"
}
