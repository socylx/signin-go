package todo

import (
	"time"
)

// ShopPromotion 促销活动集合
type ShopPromotion struct {
	ID          uint32    `gorm:"primaryKey;column:id" json:"id"`
	Name        string    `gorm:"column:name" json:"name"`               // 促销名称
	Description string    `gorm:"column:description" json:"description"` // 活动描述
	VideoURL    string    `gorm:"column:video_url" json:"video_url"`     // 视频链接
	Img         string    `gorm:"column:img" json:"img"`                 // 活动图片
	Deadline    time.Time `gorm:"column:deadline" json:"deadline"`       // 活动结束时间
	IsDel       uint8     `gorm:"column:is_del" json:"is_del"`
	CreateTime  time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime  time.Time `gorm:"column:update_time" json:"update_time"`
}

// TableName get sql table name.获取数据库表名
func (m *ShopPromotion) TableName() string {
	return "shop_promotion"
}
