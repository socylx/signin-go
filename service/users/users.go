package users

var UsersService *service

func init() {
	UsersService = &service{}
}

type Service interface {
	i()
}

type service struct{}

func (s *service) i() {}
