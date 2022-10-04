package membership

import (
	"signin-go/global/mysql"
	"signin-go/global/time"
	"signin-go/internal/core"
	"signin-go/repository/card"
)

type MembershipFilter struct {
	CreateTimeGE   time.Time
	CreateTimeLT   time.Time
	CardID         card.CardID
	IncludeUserIDs []uint32
}

type MembershipData struct {
	ID              uint32
	UserID          uint32
	BelongsStudioID uint32
	Amount          float32
	SalesUserID     uint32
}

/*
获取会员卡
*/
func GetMembershipDatas(ctx core.StdContext, filter *MembershipFilter) (data []*MembershipData, err error) {
	db := mysql.DB.WithContext(ctx)

	query := db.Table(tableName()).
		Select("membership.id,membership.user_id,membership.belongs_studio_id,membership.amount,membership.sales_user_id").
		Where("membership.is_del = 0")
	if filter.CreateTimeGE != time.TimeZeroTime {
		query = query.Where("membership.create_time >= ?", filter.CreateTimeGE)
	}
	if filter.CreateTimeLT != time.TimeZeroTime {
		query = query.Where("membership.create_time < ?", filter.CreateTimeLT)
	}
	if filter.CardID > 0 {
		query = query.Where("membership.card_id = ?", filter.CardID)
	}
	if len(filter.IncludeUserIDs) > 0 {
		query = query.Where("membership.user_id IN ?", filter.IncludeUserIDs)
	}

	err = query.Find(&data).Error
	return
}
