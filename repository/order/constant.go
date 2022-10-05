package order

type OrderStatus = int16

const (
	STATUS_CART              OrderStatus = 1
	STATUS_ORDER_IMMEDIATELY OrderStatus = 2
	STATUS_WAIT_PAY          OrderStatus = 3  // 待支付
	STATUS_WAIT_SEND         OrderStatus = 4  // 已支付
	STATUS_WAIT_RECEIVE      OrderStatus = 5  // 待收货，已发货
	STATUS_COMPLETE          OrderStatus = 6  // 订单完成（已入场，门票失效）
	STATUS_CLOSE             OrderStatus = 7  // 关闭（超时未支付等）
	STATUS_REFUND_CLOSE      OrderStatus = 8  // 退款关闭
	STATUS_MANUAL_CLOSE      OrderStatus = 9  // 手动关闭的（退款）
	STATUS_USER_CLOSE_ING    OrderStatus = 10 // 用户取消订单中
	STATUS_USER_CLOSE        OrderStatus = 11 // 用户取消订单成功
)
