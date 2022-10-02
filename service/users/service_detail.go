package users

import (
	"signin-go/internal/core"
	"signin-go/repository/users"
)

func (s *service) List(ctx core.Context, userID uint32) (user []*users.Users, err error) {
	return s.usersRepo.List(ctx, userID)
}
