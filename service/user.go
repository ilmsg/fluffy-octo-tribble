package service

import "github.com/ilmsg/fluffy-octo-tribble/domain"

type UserService struct {
	repo domain.IUserRepository
}

func NewUserService(repo domain.IUserRepository) domain.IUserService {
	return &UserService{repo}
}
