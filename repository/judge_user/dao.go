package judge_user

import (
	"gsteps-go/global/mysql"
	"gsteps-go/internal/core"
)

func GetJudgeUserDatas(ctx core.StdContext, userID uint32) (data []*JudgeUserData, err error) {
	db := mysql.DB.WithContext(ctx)
	err = db.Table("judge_user").
		Select("judge_user.id").
		Where("judge_user.is_del = 0 AND judge_user.user_id = ?", userID).
		Find(&data).Error
	return
}
