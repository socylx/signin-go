package todo

import (
	"time"
)

// ShopAddressProvince 省
type ShopAddressProvince struct {
	ID         uint32    `gorm:"primaryKey;column:id" json:"id"`
	Name       string    `gorm:"column:name" json:"name"`
	ProvinceID string    `gorm:"column:province_id" json:"province_id"`
	IsDel      bool      `gorm:"column:is_del" json:"is_del"`
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"`
}

// TableName get sql table name.获取数据库表名
func (m *ShopAddressProvince) TableName() string {
	return "shop_address_province"
}
