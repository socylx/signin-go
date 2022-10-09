package users

import (
	"errors"
	"signin-go/global/mysql"
	"signin-go/internal/core"

	"go.uber.org/zap"
)

func Detail(ctx core.StdContext, userID uint32) (user *Users, err error) {
	// db := mysql.DB.WithContext(ctx)
	user = &Users{}
	err = ctx.DB.Table("users").
		Where("users.is_del = 0 AND users.id = ?", userID).
		First(user).Error
	// ctx.Logger.Error("users.Detail", zap.Uint32("userID", userID))
	ctx.Error("users.Detail", zap.Error(errors.New("userID")))
	return
}

type Filter struct {
	ID   uint32
	Page int
	Size int
}

func List(ctx core.StdContext, filter *Filter) (users []*Users, err error) {
	db := mysql.DB.WithContext(ctx)

	sql := db.Table("users").Where("users.is_del = 0")
	if filter.ID > 0 {
		sql = sql.Where("users.id = ?", filter.ID)
	}

	if filter.Page <= 0 {
		filter.Page = 1
	}
	if filter.Size <= 0 {
		filter.Size = 20
	}

	err = sql.Limit(filter.Size).Offset(filter.Page*filter.Size - filter.Size).Find(&users).Error
	return
}

func Update(ctx core.StdContext, filter *Filter, user Users) (err error) {
	db := mysql.DB.WithContext(ctx)
	err = db.Table("users").
		Where("users.is_del = 0 AND users.id = ?", filter.ID).
		Updates(user).Error
	return
}
