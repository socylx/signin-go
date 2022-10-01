package todo

import (
	"time"
)

// ShopBrandResources [...]
type ShopBrandResources struct {
	ID          uint32    `gorm:"primaryKey;column:id" json:"id"`
	ShopBrandID uint32    `gorm:"column:shop_brand_id" json:"shop_brand_id"` // 商品品牌ID
	URL         string    `gorm:"column:url" json:"url"`                     // 资源链接
	SourceType  uint32    `gorm:"column:source_type" json:"source_type"`     // 资源类型,1:图片,2:视频，3:音频
	CreateTime  time.Time `gorm:"column:create_time" json:"create_time"`     // 创建时间
	UpdateTime  time.Time `gorm:"column:update_time" json:"update_time"`     // 更新时间
	IsDel       bool      `gorm:"column:is_del" json:"is_del"`               // 是否删除
}

// TableName get sql table name.获取数据库表名
func (m *ShopBrandResources) TableName() string {
	return "shop_brand_resources"
}
