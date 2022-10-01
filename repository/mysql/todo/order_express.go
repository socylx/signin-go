package todo

import (
	"time"
)

// OrderExpress [...]
type OrderExpress struct {
	ID               uint32    `gorm:"primaryKey;column:id" json:"id"`
	OrderID          uint32    `gorm:"column:order_id" json:"order_id"`                     // 订单ID
	AreaID           uint32    `gorm:"column:area_id" json:"area_id"`                       // 区域ID
	ExpressCode      string    `gorm:"column:express_code" json:"express_code"`             // 快递单号
	ShippingOrderNo  string    `gorm:"column:shipping_order_no" json:"shipping_order_no"`   // 运输订单号, 用于向快递公司下单，下单所用的订单号不可重复
	ExpressCompanyID uint32    `gorm:"column:express_company_id" json:"express_company_id"` // 快递公司ID
	ExpressTypeID    uint32    `gorm:"column:express_type_id" json:"express_type_id"`       // 快递运输方式ID
	WaybillURL       string    `gorm:"column:waybill_url" json:"waybill_url"`               // 运货单
	Status           uint32    `gorm:"column:status" json:"status"`                         // 快递状态
	Fee              float32   `gorm:"column:fee" json:"fee"`                               // 运费
	PayMethod        uint32    `gorm:"column:pay_method" json:"pay_method"`                 // 付款方式, 1:寄付月结, 2: 寄付现结, 3:到付
	Remark           string    `gorm:"column:remark" json:"remark"`                         // 备注
	SysRemark        string    `gorm:"column:sys_remark" json:"sys_remark"`                 // 系统备注
	OptUserID        uint32    `gorm:"column:opt_user_id" json:"opt_user_id"`               // 发货人
	CreateTime       time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime       time.Time `gorm:"column:update_time" json:"update_time"`
	IsDel            bool      `gorm:"column:is_del" json:"is_del"`
}

// TableName get sql table name.获取数据库表名
func (m *OrderExpress) TableName() string {
	return "order_express"
}
