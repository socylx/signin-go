package show_video_user

import (
	"gsteps-go/global/mysql"
	"gsteps-go/internal/core"
)

type Filter struct {
	Type     ShowVideoUserType
	PlayerID uint32
}

/*
获取学员写有记录参与的作品IDs
*/
func GetShowVideoIDs(ctx core.StdContext, filter *Filter) (ids []uint32, err error) {
	db := mysql.DB.WithContext(ctx)
	query := db.Table("show_video_user").Select("show_video_user.show_video_id").Where("show_video_user.is_del = 0")
	if filter.Type > 0 {
		query = query.Where("show_video_user.type = ?", filter.Type)
	}
	if filter.PlayerID > 0 {
		query = query.Where("show_video_user.player_id = ?", filter.PlayerID)
	}
	err = query.Find(&ids).Error
	return
}
