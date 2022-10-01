package todo

import (
	"time"
)

// ShopAddress 收货地址
type ShopAddress struct {
	ID         uint32    `gorm:"primaryKey;column:id" json:"id"`
	UserID     uint32    `gorm:"column:user_id" json:"user_id"`         // 地址所属的用户id
	ProvinceID uint32    `gorm:"column:province_id" json:"province_id"` // 省份id
	CityID     uint32    `gorm:"column:city_id" json:"city_id"`         // 市id
	CountyID   int       `gorm:"column:county_id" json:"county_id"`     // 区id
	StreetID   uint32    `gorm:"column:street_id" json:"street_id"`     // 街道id
	Name       string    `gorm:"column:name" json:"name"`               // 收货人姓名
	Detail     string    `gorm:"column:detail" json:"detail"`           // 详细地址
	Phone      string    `gorm:"column:phone" json:"phone"`             // 联系电话
	IsDefault  bool      `gorm:"column:is_default" json:"is_default"`   // 默认地址
	IsDel      bool      `gorm:"column:is_del" json:"is_del"`
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"`
}

// TableName get sql table name.获取数据库表名
func (m *ShopAddress) TableName() string {
	return "shop_address"
}
