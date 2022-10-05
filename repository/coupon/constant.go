package coupon

type CouponType = int16

const (
	OffType      CouponType = 1 //满减券
	CashType     CouponType = 2 //代金券
	DiscountType CouponType = 3 //折扣券
)

type CouponAmountType = int16

const (
	CashAmountType   CouponAmountType = 1 //金额
	NumberAmountType CouponAmountType = 2 //次数
	TimeAmountType   CouponAmountType = 3 //时长
)
