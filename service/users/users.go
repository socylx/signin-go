package users

import (
	"signin-go/internal/core"
	"signin-go/repository/users"
)

type Service interface {
	i()

	Detail(ctx core.Context, userID uint32) (user *users.Users, err error)
	List(ctx core.Context, userID uint32) (user []*users.Users, err error)
	Update(ctx core.Context, userID uint32) (err error)
}

type service struct {
	usersRepo users.UsersRepo
}

func New() Service {
	return &service{
		usersRepo: users.New(),
	}
}

func (s *service) i() {}
