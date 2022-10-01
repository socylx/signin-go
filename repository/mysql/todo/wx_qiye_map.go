package todo

import (
	"time"
)

// WxQiyeMap [...]
type WxQiyeMap struct {
	ID         uint32    `gorm:"primaryKey;column:id" json:"id"`
	QiyeUserid string    `gorm:"column:qiye_userid" json:"qiye_userid"`
	ForID      uint32    `gorm:"column:for_id" json:"for_id"`
	ForType    uint32    `gorm:"column:for_type" json:"for_type"`
	RoleType   uint32    `gorm:"column:role_type" json:"role_type"` // 1-staff, 2-friend
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"`
	IsDel      bool      `gorm:"column:is_del" json:"is_del"`
}

// TableName get sql table name.获取数据库表名
func (m *WxQiyeMap) TableName() string {
	return "wx_qiye_map"
}
