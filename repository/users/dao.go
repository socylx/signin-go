package users

func (u *users) Detail(userID uint32) (user *Users, err error) {
	user = &Users{}

	err = u.db.Table(
		u.TableName(),
	).Where(
		"users.is_del = 0 AND users.id = ?", userID,
	).First(user).Error
	return user, err
}
