package todo

import (
	"time"
)

// StudioResources 门店的图片，视频资源
type StudioResources struct {
	ID         uint32    `gorm:"primaryKey;column:id" json:"id"`
	URL        string    `gorm:"column:url" json:"url"`           // 资源链接
	Title      string    `gorm:"column:title" json:"title"`       // 资源描述
	Type       uint32    `gorm:"column:type" json:"type"`         // 资源类型，1图片，2视频，3音频
	ForType    uint32    `gorm:"column:for_type" json:"for_type"` // 资源所属，1门店地址资源, 2门店介绍资源
	ForID      uint32    `gorm:"column:for_id" json:"for_id"`     // 所属资源id
	Index      uint32    `gorm:"column:index" json:"index"`       // 资源显示排序，越小越排在前边
	IsDel      bool      `gorm:"column:is_del" json:"is_del"`
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"`
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
	OptUserID  uint32    `gorm:"column:opt_user_id" json:"opt_user_id"` // 上传人id
}

// TableName get sql table name.获取数据库表名
func (m *StudioResources) TableName() string {
	return "studio_resources"
}
