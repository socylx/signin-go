package show_video

import (
	"gsteps-go/global/mysql"
	"gsteps-go/internal/core"
	"gsteps-go/repository/show_video_user"
	"gsteps-go/repository/signin"
)

/*
获取学员参与作品的数量
*/
func GetShowVideoCount(ctx core.StdContext, userID uint32) (count int64, err error) {
	db := mysql.DB.WithContext(ctx)
	showVideoIDs, _ := show_video_user.GetShowVideoIDs(ctx, &show_video_user.Filter{
		Type:     show_video_user.User,
		PlayerID: userID,
	})
	err = db.Table("show_video").
		Select("show_video.id").
		Joins("JOIN signin ON show_video.activity_id = signin.activity_id").
		Where("show_video.is_del = 0").
		Where("signin.is_del = 0 AND signin.is_back = 0 AND signin.type IN ? AND signin.user_id = ?", signin.StatusOnList, userID).
		Or("show_video.id in ?", showVideoIDs).
		Count(&count).Error
	return

}
