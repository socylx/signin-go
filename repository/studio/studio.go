package studio

import (
	"time"
)

// Studio [...]
type Studio struct {
	ID         uint32    `gorm:"primaryKey;column:id" json:"id"`
	BrandID    int       `gorm:"column:brand_id" json:"brand_id"`       // 品牌id
	Phone      string    `gorm:"column:phone" json:"phone"`             // 联系人user_id
	Province   string    `gorm:"column:province" json:"province"`       // 省
	City       string    `gorm:"column:city" json:"city"`               // 市
	Zone       string    `gorm:"column:zone" json:"zone"`               // 区
	Address    string    `gorm:"column:address" json:"address"`         // 详细地址
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"` // 创建时间
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"` // 更新时间
	IsDel      bool      `gorm:"column:is_del" json:"is_del"`           // 是否被删除
	Name       string    `gorm:"column:name" json:"name"`               // 场馆名称
	Desc       string    `gorm:"column:desc" json:"desc"`               // 场馆介绍
	LogoURL    string    `gorm:"column:logo_url" json:"logo_url"`       // 场馆logo
	Latitude   string    `gorm:"column:latitude" json:"latitude"`       // 门店位置纬度
	Longitude  string    `gorm:"column:longitude" json:"longitude"`     // 门店位置经度
	StudioType uint32    `gorm:"column:studio_type" json:"studio_type"` // 门店类型，1线下门店，2虚拟门店
}
