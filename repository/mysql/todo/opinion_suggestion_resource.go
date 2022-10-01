package todo

import (
	"time"
)

// OpinionSuggestionResource [...]
type OpinionSuggestionResource struct {
	ID                  uint32    `gorm:"primaryKey;column:id" json:"id"`
	OpinionSuggestionID uint32    `gorm:"column:opinion_suggestion_id" json:"opinion_suggestion_id"` // 反馈的id
	ForType             uint32    `gorm:"column:for_type" json:"for_type"`                           // 1、学员反馈的相关资源, 2、工作人员处理时相关资源
	URL                 string    `gorm:"column:url" json:"url"`                                     // 资源链接
	ResourceType        uint32    `gorm:"column:resource_type" json:"resource_type"`                 // 资源类型,1:图片,2:视频，3:音频
	CreateTime          time.Time `gorm:"column:create_time" json:"create_time"`                     // 创建时间
	UpdateTime          time.Time `gorm:"column:update_time" json:"update_time"`                     // 更新时间
	IsDel               bool      `gorm:"column:is_del" json:"is_del"`                               // 是否删除
}

// TableName get sql table name.获取数据库表名
func (m *OpinionSuggestionResource) TableName() string {
	return "opinion_suggestion_resource"
}
