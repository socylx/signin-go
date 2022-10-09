package user_snapshot

import (
	"gsteps-go/global/mysql"
	"gsteps-go/global/time"
	"gsteps-go/internal/core"
)

type UserSnapshotData struct {
	ID                uint32
	MembershipRemains float32
	CouponRemains     float32
	RenewMembershipID uint32
	UserID            uint32
}
type UserSnapshotFilter struct {
	CreateTimeGE time.Time
	CreateTimeLT time.Time
	StudioID     uint32
	IncludeIDs   []uint32
}

/*
获取某个门店的学员在某段时间内的节点数据
*/
func GetUserSnapshotData(ctx core.StdContext, filter *UserSnapshotFilter) (data []*UserSnapshotData, err error) {
	db := mysql.DB.WithContext(ctx)
	query := db.Table("user_snapshot").
		Select("user_snapshot.id,user_snapshot.membership_remains,user_snapshot.coupon_remains,user_snapshot.renew_membership_id,users.id as user_id").
		Joins("JOIN users on user_snapshot.user_id = users.id").
		Where("user_snapshot.is_del = 0 AND user_snapshot.user_type = 1 AND users.is_del = 0")

	if filter.CreateTimeGE != time.TimeZeroTime {
		query = query.Where("user_snapshot.create_time >= ?", filter.CreateTimeGE)
	}
	if filter.CreateTimeLT != time.TimeZeroTime {
		query = query.Where("user_snapshot.create_time < ?", filter.CreateTimeLT)
	}
	if filter.StudioID > 0 {
		query = query.Where("users.belongs_studio_id = ?", filter.StudioID)
	}
	if len(filter.IncludeIDs) > 0 {
		query = query.Where("user_snapshot.id IN ?", filter.IncludeIDs)
	}
	err = query.Find(&data).Error
	return
}

type LastUserSnapshotIDsFilter struct {
	Time           time.Time
	StudioID       uint32
	ExcludeUserIDs []uint32
}

/*
获取学员最后一个节点的ID
*/
func GetLastUserSnapshotIDs(ctx core.StdContext, filter *LastUserSnapshotIDsFilter) (data []uint32, err error) {
	db := mysql.DB.WithContext(ctx)

	query := db.Table("user_snapshot").
		Select("max(user_snapshot.id) as last_user_snapshot_id").
		Joins("JOIN users on user_snapshot.user_id = users.id").
		Where("user_snapshot.is_del = 0 AND user_snapshot.create_time < ? AND user_snapshot.user_type = 1 AND users.is_del = 0 AND users.belongs_studio_id = ?", filter.Time, filter.StudioID)
	if len(filter.ExcludeUserIDs) > 0 {
		query = query.Where("user_snapshot.user_id NOT IN ?", filter.ExcludeUserIDs)
	}
	err = query.Group("user_snapshot.user_id").Find(&data).Error
	return
}
