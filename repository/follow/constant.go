package follow

type FollowType uint
type ForType uint

const (
	BeforeSold        FollowType = 1 //售前(转介绍，新人)
	Renewal           FollowType = 2 //售后(续费)
	MembershipService FollowType = 3 //售后(会员服务)

	User             ForType = 1 //这条根据记录是属于注册用户的
	UserBeforeMember ForType = 2 //这条根据记录是属于线索的
)
