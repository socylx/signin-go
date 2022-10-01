package todo

import (
	"time"
)

// UserOther 用户关联的三方号的信息，例如某个用户关注了公众号，对应的opendid 会存到这个表
type UserOther struct {
	ID           uint32    `gorm:"primaryKey;column:id" json:"id"`
	UserID       uint32    `gorm:"column:user_id" json:"user_id"` // 用户id
	UserBeforeID uint32    `gorm:"column:user_before_id" json:"user_before_id"`
	Type         uint32    `gorm:"column:type" json:"type"`         // 账号类型，1G社服务号openid
	OpenID       string    `gorm:"column:open_id" json:"open_id"`   // 其他系统的用户id, 例如微信公众号的openid
	UnionID      string    `gorm:"column:union_id" json:"union_id"` // 平台union_id
	IsDel        bool      `gorm:"column:is_del" json:"is_del"`
	CreateTime   time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime   time.Time `gorm:"column:update_time" json:"update_time"`
}

// TableName get sql table name.获取数据库表名
func (m *UserOther) TableName() string {
	return "user_other"
}
