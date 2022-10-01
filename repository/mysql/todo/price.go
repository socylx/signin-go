package todo

import (
	"time"
)

// Price [...]
type Price struct {
	ID         uint32    `gorm:"primaryKey;column:id" json:"id"`
	Limit      int       `gorm:"column:limit" json:"limit"`             // 额度
	Price      float32   `gorm:"column:price" json:"price"`             // 理论收款
	Expire     int       `gorm:"column:expire" json:"expire"`           // 过期时间 天
	Remark     string    `gorm:"column:remark" json:"remark"`           // 说明
	Support    string    `gorm:"column:support" json:"support"`         // 支持类型, 1充值，2新购
	OnlyStaff  bool      `gorm:"column:only_staff" json:"only_staff"`   // 是否仅工作人员可见
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"` // 创建时间
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"` // 更新时间
	IsDel      bool      `gorm:"column:is_del" json:"is_del"`           // 是否已被删除
	CardID     int       `gorm:"column:card_id" json:"card_id"`         // 卡类型id
	Remains    uint32    `gorm:"column:remains" json:"remains"`         // 商品剩余数量
	SoldNum    uint32    `gorm:"column:sold_num" json:"sold_num"`       // 商品已经售出多少件
	Name       string    `gorm:"column:name" json:"name"`               // 对外显示的文案
	MaxNum     uint32    `gorm:"column:max_num" json:"max_num"`         // 每个人的最多购买数量
	OnSale     bool      `gorm:"column:on_sale" json:"on_sale"`         // 是否在售，1在，0否
}

// TableName get sql table name.获取数据库表名
func (m *Price) TableName() string {
	return "price"
}
