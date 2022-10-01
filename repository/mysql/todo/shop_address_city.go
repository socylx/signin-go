package todo

import (
	"time"
)

// ShopAddressCity 收货地址 市
type ShopAddressCity struct {
	ID         uint32    `gorm:"primaryKey;column:id" json:"id"`
	Name       string    `gorm:"column:name" json:"name"`
	ProvinceID string    `gorm:"column:province_id" json:"province_id"` // 省编号
	CityID     string    `gorm:"column:city_id" json:"city_id"`         // 市编号
	IsDel      bool      `gorm:"column:is_del" json:"is_del"`
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"`
}

// TableName get sql table name.获取数据库表名
func (m *ShopAddressCity) TableName() string {
	return "shop_address_city"
}
