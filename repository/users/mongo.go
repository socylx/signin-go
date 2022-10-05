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
	PageAccessData   *PageAccessData               `bson:"page_access_data"`
	PageEventData    *PageEventData                `bson:"page_event_data"`
	ShowVideoCount   int64                         `bson:"show_video_count"`
	AllSigninSpend   float64                       `bson:"all_signin_spend"`
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

type PageAccessData struct {
	PageAccessCount                  uint64    `bson:"page_access_count" json:"page_access_count"`
	LastPageAccessTime               time.Time `bson:"last_page_access_time" json:"last_page_access_time"`
	CurrentStudioAccessActivityCount uint64    `bson:"current_studio_access_activity_count" json:"current_studio_access_activity_count"`
	AccessBuyCardCount               uint64    `bson:"access_buy_card_count" json:"access_buy_card_count"`
}

type PageEventData struct {
	AccessLocation *AccessLocation `bson:"access_location" json:"access_location"`
}

type AccessLocation struct {
	Longitude string `bson:"longitude" json:"longitude"`
	Latitude  string `bson:"latitude" json:"latitude"`
}
