package users

import (
	"signin-go/internal/core"
	"signin-go/repository/users"
)

func (s *service) Detail(ctx core.Context, userID uint32) (user *users.Users, err error) {
	return s.usersRepo.Detail(ctx, 3352)
}

func (s *service) List(ctx core.Context, userID uint32) (user []*users.Users, err error) {
	return s.usersRepo.List(ctx, &users.Filter{ID: 3352})
}

func (s *service) Update(ctx core.Context, userID uint32) (err error) {
	return s.usersRepo.Update(ctx, &users.Filter{ID: 3352}, map[string]interface{}{"users.name": "佳哥"})
}
