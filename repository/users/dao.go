package users

import (
	"signin-go/global/mysql"
	"signin-go/internal/core"
)

func Detail(ctx core.Context, userID uint32) (user *Users, err error) {
	db := mysql.DB.WithContext(ctx.RequestContext())

	user = &Users{}
	err = db.Table(
		tableName(),
	).Where(
		"users.is_del = 0 AND users.id = ?", userID,
	).First(user).Error
	return
}

type Filter struct {
	ID   uint32
	Page int
	Size int
}

func List(ctx core.Context, filter *Filter) (users []*Users, err error) {
	db := mysql.DB.WithContext(ctx.RequestContext())

	sql := db.Table(
		tableName(),
	).Where(
		"users.is_del = 0",
	)
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

func Update(ctx core.Context, filter *Filter, user Users) (err error) {
	db := mysql.DB.WithContext(ctx.RequestContext())
	err = db.Table(
		tableName(),
	).Where(
		"users.is_del = 0 AND users.id = ?", filter.ID,
	).Updates(user).Error
	return
}
