package staff

import (
	"time"
)

// Staff [...]
type Staff struct {
	ID         uint32    `gorm:"primaryKey;column:id" json:"id"`
	RoleID     int       `gorm:"column:role_id" json:"role_id"`         // 角色id
	UserID     int       `gorm:"column:user_id" json:"user_id"`         // 用户id
	StudioID   int       `gorm:"column:studio_id" json:"studio_id"`     // 所属工作室id
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"` // 更新时间
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"` // 创建时间
	IsDel      bool      `gorm:"column:is_del" json:"is_del"`           // 是否删除
	Type       uint32    `gorm:"column:type" json:"type"`               // 账号类型，1公司账号，2私人账号
	WxQrCode   string    `gorm:"column:wx_qr_code" json:"wx_qr_code"`   // 二维码链接
}

// TableName get sql table name.获取数据库表名
func tableName() string {
	return "staff"
}
