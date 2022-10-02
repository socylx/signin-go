package users

import "signin-go/internal/core"

func (u *users) List(ctx core.Context, userID uint32) (user []*Users, err error) {
	user = []*Users{}

	db := u.db.WithContext(ctx.RequestContext())
	err = db.Table(
		u.TableName(),
	).Where(
		"users.is_del = 0 AND users.id = ?", userID,
	).Find(&user).Error
	return user, err
}
