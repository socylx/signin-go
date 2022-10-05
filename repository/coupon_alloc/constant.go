package coupon_alloc

type CouponAllocStatus uint

const (
	Init        CouponAllocStatus = 0 //未使用
	Used        CouponAllocStatus = 1 //已经使用
	Invalid                       = 2 //失效
	NoActivated                   = 3 //未激活
)
