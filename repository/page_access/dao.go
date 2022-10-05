package page_access

import (
	"signin-go/global/mysql"
	"signin-go/global/time"
	"signin-go/internal/core"
)

type Filter struct {
	UserID       uint32
	CreateTimeGE time.Time
	CreateTimeLT time.Time
	Type         PageAccessType
}

func GetPageAccess(ctx core.StdContext, filter *Filter) (data []*PageAccess, err error) {
	db := mysql.DB.WithContext(ctx)
	query := db.Table("page_access").Where("page_access.is_del = 0")
	if filter.UserID > 0 {
		query = query.Where("page_access.user_id = ?", filter.UserID)
	}
	if filter.CreateTimeGE != time.TimeZeroTime {
		query = query.Where("page_access.create_time >= ?", filter.CreateTimeGE)
	}
	if filter.CreateTimeLT != time.TimeZeroTime {
		query = query.Where("page_access.create_time < ?", filter.CreateTimeLT)
	}
	if filter.Type != "" {
		query = query.Where("page_access.type = ?", filter.Type)
	}
	err = query.Find(&data).Error
	return
}

func GetLastPageAccessTime(ctx core.StdContext, filter *Filter) (t time.Time, err error) {
	db := mysql.DB.WithContext(ctx)

	pageAccess := &PageAccess{}
	err = db.Table("page_access").
		Where("page_access.is_del = 0 AND page_access.user_id = ?", filter.UserID).
		Order("page_access.create_time DESC").First(pageAccess).Error
	t = pageAccess.CreateTime
	return
}
