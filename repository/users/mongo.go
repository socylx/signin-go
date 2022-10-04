package users

import (
	"signin-go/global/time"
	"signin-go/repository/follow"
	"signin-go/repository/membership"
)

type Data struct {
	UserBeforeMember *UserBeforeMember            `bson:"user_before_member"`
	User             *User                        `bson:"user"`
	Memberships      []*membership.MembershipData `bson:"memberships"`
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
