package page_event

import (
	"gsteps-go/global/mysql"
	"gsteps-go/internal/core"
	"gsteps-go/repository/page_access"
)

type Filter struct {
	UserID   uint32
	Type     page_access.PageAccessType
	EventKey EventKey
}

func GetLastPageEvent(ctx core.StdContext, filter *Filter) (data *PageEvent, err error) {
	db := mysql.DB.WithContext(ctx)
	data = &PageEvent{}
	err = db.Table("page_event").
		Where("page_event.is_del = 0 AND page_event.user_id = ? AND page_event.type = ? AND page_event.event_key = ?", filter.UserID, filter.Type, filter.EventKey).
		Order("page_event.id DESC").First(data).Error
	return
}
