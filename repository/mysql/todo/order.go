package todo

import (
	"time"
)

// Order [...]
type Order struct {
	ID                uint32    `gorm:"primaryKey;column:id" json:"id"`
	UserID            uint32    `gorm:"column:user_id" json:"user_id"`
	Status            int16     `gorm:"column:status" json:"status"`               // 3未付款，4购买成功，6订单完成
	OrderNo           string    `gorm:"column:order_no" json:"order_no"`           // 微信生成的订单号
	OutTradeNo        string    `gorm:"column:out_trade_no" json:"out_trade_no"`   // 商户系统内部订单号
	OutRefundNo       string    `gorm:"column:out_refund_no" json:"out_refund_no"` // 退款订单号
	Total             float32   `gorm:"column:total" json:"total"`
	Discount          float32   `gorm:"column:discount" json:"discount"` // 折扣多少钱
	IsDel             bool      `gorm:"column:is_del" json:"is_del"`
	PayTime           time.Time `gorm:"column:pay_time" json:"pay_time"`
	SendTime          time.Time `gorm:"column:send_time" json:"send_time"` // 发货时间
	DoneTime          time.Time `gorm:"column:done_time" json:"done_time"` // 订单完成时间
	UpdateTime        time.Time `gorm:"column:update_time" json:"update_time"`
	CreateTime        time.Time `gorm:"column:create_time" json:"create_time"`
	OptUserID         uint32    `gorm:"column:opt_user_id" json:"opt_user_id"`                 // 操作人员 id, 修改订单状态时记录操作人
	CouponID          uint32    `gorm:"column:coupon_id" json:"coupon_id"`                     // 订单使用的券id
	CouponAllocID     uint32    `gorm:"column:coupon_alloc_id" json:"coupon_alloc_id"`         // 订单使用的券实例id
	PaymentType       int       `gorm:"column:payment_type" json:"payment_type"`               // 支付方式，1微信支付，2支付宝
	AddressUsername   string    `gorm:"column:address_username" json:"address_username"`       // 收货人姓名
	AddressPhone      string    `gorm:"column:address_phone" json:"address_phone"`             // 收货人手机号
	AddressDetail     string    `gorm:"column:address_detail" json:"address_detail"`           // 收货地址
	ProvinceID        uint32    `gorm:"column:province_id" json:"province_id"`                 // 收货地址 省id
	CityID            uint32    `gorm:"column:city_id" json:"city_id"`                         // 收货地址 市id
	CountyID          uint32    `gorm:"column:county_id" json:"county_id"`                     // 收货地址 区id
	ExpressCode       string    `gorm:"column:express_code" json:"express_code"`               // 快递单号
	ShippingOrderNo   string    `gorm:"column:shipping_order_no" json:"shipping_order_no"`     // 运输订单号, 用于向快递公司下单，下单所用的订单号不可重复
	ExpressDeliveryID string    `gorm:"column:express_delivery_id" json:"express_delivery_id"` // 快递公司编号：SF顺丰，YTO圆通，EMS邮政
	ExpressCompanyID  uint32    `gorm:"column:express_company_id" json:"express_company_id"`   // 快递公司ID
	ExpressTypeID     uint32    `gorm:"column:express_type_id" json:"express_type_id"`         // 快递运输方式ID
	ExpressFee        float32   `gorm:"column:express_fee" json:"express_fee"`                 // 支付的运费
	WaybillURL        string    `gorm:"column:waybill_url" json:"waybill_url"`                 // 运货单
	ShopGroupAllocID  uint32    `gorm:"column:shop_group_alloc_id" json:"shop_group_alloc_id"` // 所属拼团实例id
	SalesUserID       uint32    `gorm:"column:sales_user_id" json:"sales_user_id"`             // 售卡人
	SalesStudioID     uint32    `gorm:"column:sales_studio_id" json:"sales_studio_id"`         // 售卡门店
}

// TableName get sql table name.获取数据库表名
func (m *Order) TableName() string {
	return "order"
}
