package membership

import (
	"gsteps-go/global/mysql"
	"gsteps-go/global/time"
	"gsteps-go/internal/core"
	"gsteps-go/repository/card"
)

type MembershipFilter struct {
	UserID         uint32
	CreateTimeGE   time.Time
	CreateTimeLT   time.Time
	CardIDs        []card.CardID
	RemainsGT      float32
	IncludeUserIDs []uint32
	Status         uint // 0-不限制, 1-有效期未过
}

type MembershipData struct {
	ID              uint32    `bson:"id"`
	UserID          uint32    `bson:"user_id"`
	BelongsStudioID uint32    `bson:"belongs_studio_id"`
	Amount          float32   `bson:"amount"`
	SalesUserID     uint32    `bson:"sales_user_id"`
	CreateTime      time.Time `bson:"create_time"`
	Deadline        time.Time `bson:"deadline"`
	Remains         float32   `bson:"remains"`
}

/*
获取会员卡
*/
func GetMembershipDatas(ctx core.StdContext, filter *MembershipFilter) (data []*MembershipData, err error) {
	db := mysql.DB.WithContext(ctx)

	query := db.Table("membership").
		Select("membership.id,membership.user_id,membership.belongs_studio_id,membership.amount,membership.sales_user_id,membership.create_time,membership.deadline,membership.remains").
		Where("membership.is_del = 0")
	if filter.UserID > 0 {
		query = query.Where("membership.user_id = ?", filter.UserID)
	}
	if filter.CreateTimeGE != time.TimeZeroTime {
		query = query.Where("membership.create_time >= ?", filter.CreateTimeGE)
	}
	if filter.CreateTimeLT != time.TimeZeroTime {
		query = query.Where("membership.create_time < ?", filter.CreateTimeLT)
	}
	if len(filter.CardIDs) > 0 {
		query = query.Where("membership.card_id IN ?", filter.CardIDs)
	}
	if filter.RemainsGT > 0 {
		query = query.Where("membership.remains > ?", filter.RemainsGT)
	}
	if len(filter.IncludeUserIDs) > 0 {
		query = query.Where("membership.user_id IN ?", filter.IncludeUserIDs)
	}
	if filter.Status == 1 {
		query = query.Where("membership.deadline = ? OR membership.deadline > ?", time.TimeZeroString, time.Now)
	}

	err = query.Find(&data).Error
	return
}

func GetMaxMembershipID(ctx core.StdContext) (count int64, err error) {
	db := mysql.DB.WithContext(ctx)
	err = db.Table("membership").
		Select("membership.id").Count(&count).Error
	return
}

type Filter struct {
	IDGT int64
	Size int
}

func GetUserBeforeMemberIDs(ctx core.StdContext, filter *Filter) (ids []int64, err error) {
	db := mysql.DB.WithContext(ctx)
	err = db.Table("membership").
		Select("membership.user_id").
		Where("membership.id > ?", filter.IDGT).
		Limit(filter.Size).
		Find(&ids).Error
	return
}
