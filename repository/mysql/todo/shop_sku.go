package todo

import (
	"time"
)

// ShopSku sku
type ShopSku struct {
	ID              uint32    `gorm:"primaryKey;column:id" json:"id"`
	SpuID           uint32    `gorm:"column:spu_id" json:"spu_id"`                       // 所属spuid
	No              uint32    `gorm:"column:no" json:"no"`                               // 编号，与spu组合成 sku的号
	Price           float64   `gorm:"column:price" json:"price"`                         // 价格
	Tips            string    `gorm:"column:tips" json:"tips"`                           // sku 商品的提示
	AttributeStr    string    `gorm:"column:attribute_str" json:"attribute_str"`         // 商品属性串，如 1:2,2:1  代表有两个属性，属性颜色(1)红色(2),尺寸(2)黑色(1)
	AttributeStrMd5 string    `gorm:"column:attribute_str_md5" json:"attribute_str_md5"` // 根据 attribute_str 算出来的MD5值
	Stock           uint32    `gorm:"column:stock" json:"stock"`                         // 库存
	OptUserID       uint32    `gorm:"column:opt_user_id" json:"opt_user_id"`             // 修改人id
	IsDel           bool      `gorm:"column:is_del" json:"is_del"`
	UpdateTime      time.Time `gorm:"column:update_time" json:"update_time"`
	CreateTime      time.Time `gorm:"column:create_time" json:"create_time"`
}

// TableName get sql table name.获取数据库表名
func (m *ShopSku) TableName() string {
	return "shop_sku"
}
