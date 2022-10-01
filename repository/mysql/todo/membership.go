package todo

import (
	"time"
)

// Membership [...]
type Membership struct {
	ID              uint32    `gorm:"primaryKey;column:id" json:"id"`
	CardID          int       `gorm:"column:card_id" json:"card_id"`               // 会员卡id
	UserID          int       `gorm:"column:user_id" json:"user_id"`               // 会员用户id
	Deadline        time.Time `gorm:"column:deadline" json:"deadline"`             // 有效期
	CreateTime      time.Time `gorm:"column:create_time" json:"create_time"`       // 创建时间
	UpdateTime      time.Time `gorm:"column:update_time" json:"update_time"`       // 更新时间
	IsDel           bool      `gorm:"column:is_del" json:"is_del"`                 // 1为删除
	Remains         float32   `gorm:"column:remains" json:"remains"`               // 余额
	Amount          float32   `gorm:"column:amount" json:"amount"`                 // 实际支付金额
	EntityCardID    string    `gorm:"column:entity_card_id" json:"entity_card_id"` // 实体卡ID
	Payment         int       `gorm:"column:payment" json:"payment"`               // 支付方式
	Remark          string    `gorm:"column:remark" json:"remark"`                 // 备注
	PriceID         int       `gorm:"column:price_id" json:"price_id"`             // 价格id
	BelongsStudioID int       `gorm:"column:belongs_studio_id" json:"belongs_studio_id"`
	SalesUserID     int       `gorm:"column:sales_user_id" json:"sales_user_id"`   // 销售人员user表id
	SysStatus       uint32    `gorm:"column:sys_status" json:"sys_status"`         // 系统标记卡状态，1正常，2退款退卡，3开错删除，4余额转为其他卡
	RefundAmount    float32   `gorm:"column:refund_amount" json:"refund_amount"`   // 退卡退款金额
	CreateUserID    uint32    `gorm:"column:create_user_id" json:"create_user_id"` // 创建人
}

// TableName get sql table name.获取数据库表名
func (m *Membership) TableName() string {
	return "membership"
}
