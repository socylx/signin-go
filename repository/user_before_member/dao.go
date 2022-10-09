package user_before_member

import (
	"gsteps-go/global/mysql"
	"gsteps-go/internal/core"
)

func GetMaxUserBeforeMemberID(ctx core.StdContext) (count int64, err error) {
	db := mysql.DB.WithContext(ctx)
	err = db.Table("user_before_member").
		Select("user_before_member.id").Count(&count).Error
	return
}

type Filter struct {
	IDGT int64
	Size int
}

func GetUserBeforeMemberIDs(ctx core.StdContext, filter *Filter) (ids []int64, err error) {
	db := mysql.DB.WithContext(ctx)
	err = db.Table("user_before_member").
		Select("user_before_member.id").
		Where("user_before_member.id > ?", filter.IDGT).
		Limit(filter.Size).
		Find(&ids).Error
	return
}

func Detail(ctx core.StdContext, ID, userID uint32) (data *UserBeforeMember, err error) {
	data = &UserBeforeMember{}
	if ID <= 0 && userID <= 0 {
		return
	}
	db := mysql.DB.WithContext(ctx)
	query := db.Table("user_before_member").
		Where("user_before_member.is_del = 0")
	if ID > 0 {
		query = query.Where("user_before_member.id = ?", ID)
	}
	if userID > 0 {
		query = query.Where("user_before_member.user_id = ?", userID)
	}
	err = query.First(data).Error
	return
}
