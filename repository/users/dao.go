package users

func (u *users) Detail(userID uint32) (user *Users) {
	user = &Users{}

	u.db.Table(
		u.TableName(),
	).Where(
		"users.is_del = 0 AND users.id = ?", userID,
	).First(user)

	return user
}
