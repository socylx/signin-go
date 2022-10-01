package todo

import (
	"time"
)

// ShopCommentResources 商品评价的资源
type ShopCommentResources struct {
	ID         uint32    `gorm:"primaryKey;column:id" json:"id"`
	URL        string    `gorm:"column:url" json:"url"`       // 资源链接
	Type       uint32    `gorm:"column:type" json:"type"`     // 资源类型，1图片，2视频，3音频
	ForID      uint32    `gorm:"column:for_id" json:"for_id"` // 资源所属id
	Index      uint32    `gorm:"column:index" json:"index"`   // 资源显示排序，越小越排在前边
	OptUserID  uint32    `gorm:"column:opt_user_id" json:"opt_user_id"`
	IsDel      bool      `gorm:"column:is_del" json:"is_del"`
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"`
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
}

// TableName get sql table name.获取数据库表名
func (m *ShopCommentResources) TableName() string {
	return "shop_comment_resources"
}
