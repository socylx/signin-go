package fission_map

type FissionType uint

const (
	Member FissionType = 1 //学员邀请学员
	Admin  FissionType = 2 //管理员邀请学员
)

type FissionStatus uint

const (
	STATUS_INIT             FissionStatus = 1 // 1邀请还未接受
	STATUS_ACCEPT           FissionStatus = 2 // 2邀请接受
	STATUS_BUY              FissionStatus = 3 // 3已购买体验券
	STATUS_ACTIVITY_FINISH  FissionStatus = 4 // 4已上课
	STATUS_ADD_MEMBERSHIP   FissionStatus = 5 // 5已经办卡
	STATUS_OLD_USER         FissionStatus = 6 // 6老用户
	STATUS_OTHER            FissionStatus = 7 // 7已经被别人邀请
	STATUS_HAS_COUPON       FissionStatus = 8 // 8邀请失败，已经有优惠券了
	STATUS_MANAGER_NEW_USER FissionStatus = 9 // 工作人员邀请的是新人
)

var (
	STATUS_AFTER_ACCEPT = []FissionStatus{STATUS_ACCEPT, STATUS_BUY, STATUS_ACTIVITY_FINISH, STATUS_ADD_MEMBERSHIP}
)
