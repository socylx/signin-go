package users

import (
	"signin-go/global/time"
	"signin-go/repository/coupon_alloc"
	"signin-go/repository/fission_map"
	"signin-go/repository/follow"
	"signin-go/repository/judge_user"
	"signin-go/repository/membership"
	"signin-go/repository/signin"
)

type Data struct {
	UserBeforeMember *UserBeforeMember             `bson:"user_before_member"`
	User             *User                         `bson:"user"`
	Memberships      []*membership.MembershipData  `bson:"memberships"`
	CouponAllocData  *CouponAllocData              `bson:"coupon_alloc_data"`
	Signins          []*signin.SigninData          `bson:"signins"`
	FissionMap       []*fission_map.FissionMapData `bson:"fission_map"`
	JudgeUserData    []*judge_user.JudgeUserData   `bson:"judge_user_data"`
}

type UserBeforeMember struct {
	ID             uint32           `bson:"id"`
	TransferTime   time.Time        `bson:"transfer_time"`
	UserID         uint32           `bson:"user_id"`
	SourceID       uint32           `bson:"source_id"`
	ManagerUserID  uint32           `bson:"manager_user_id"`
	BelongStudioID uint32           `bson:"belong_studio_id"`
	Follows        []*follow.Follow `bson:"follows"`
}

type User struct {
	ID              uint32 `bson:"id"`
	BelongsStudioID uint32 `bson:"belongs_studio_id"`
	ManagerUserID   uint32 `bson:"manager_user_id"`
}

type CouponAllocData struct {
	CouponAllocs            []*coupon_alloc.CouponAllocData `bson:"coupon_allocs"`
	LastNewUserCouponSignin *signin.SigninData              `bson:"last_new_user_coupon_signin"`
}
