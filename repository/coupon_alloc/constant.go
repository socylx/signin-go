package coupon_alloc

type CouponAllocStatus uint

const (
	Init        CouponAllocStatus = 0 //未使用
	Used        CouponAllocStatus = 1 //已经使用
	Invalid                       = 2 //失效
	NoActivated                   = 3 //未激活
)

type CouponAllocGetType = uint32

const (
	GET_TYPE_BUY            CouponAllocGetType = 1 // 订单购买
	GET_TYPE_ADMIN          CouponAllocGetType = 2 // 管理员发放
	GET_TYPE_FREE           CouponAllocGetType = 3 // 免费领取
	GET_TYPE_LAXIN          CouponAllocGetType = 4 // 老学员拉新赠送-免费, 其他都是付费(只针对新人体验券)
	GET_TYPE_LAXIN_DISCOUNT CouponAllocGetType = 5 // 老带新每月折扣活动
	GET_TYPE_LAXIN_OF_OLD   CouponAllocGetType = 6 // 老学员拉新赠送, 属于老会员
	GET_TYPE_LAXIN_OF_NEW   CouponAllocGetType = 7 // 老学员拉新赠送, 属于新会员
)
