package todo

import (
	"time"
)

// ShopAddressCounty 地址区
type ShopAddressCounty struct {
	ID         uint32    `gorm:"primaryKey;column:id" json:"id"`
	Name       string    `gorm:"column:name" json:"name"`
	CountyID   string    `gorm:"column:county_id" json:"county_id"` // 区编号
	CityID     string    `gorm:"column:city_id" json:"city_id"`     // 市编号
	IsDel      bool      `gorm:"column:is_del" json:"is_del"`
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"`
}

// TableName get sql table name.获取数据库表名
func (m *ShopAddressCounty) TableName() string {
	return "shop_address_county"
}
