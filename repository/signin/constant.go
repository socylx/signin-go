package signin

type SigninType uint

const (
	SigninStatus    SigninType = 1 //上课签到
	CancelStatus    SigninType = 2 //取消签到、预约
	SubscribeStatus SigninType = 3 //课程预约
)

var StatusOnList = [2]SigninType{SigninStatus, SubscribeStatus}
