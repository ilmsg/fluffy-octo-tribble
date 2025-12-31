package service

import (
	"github.com/ilmsg/fluffy-octo-tribble/domain"
	"github.com/ilmsg/fluffy-octo-tribble/model"
)

type UserService struct {
	repo domain.IUserRepository
}

// Create implements [domain.IUserService].
func (s *UserService) Create(user *model.User) error {
	return s.repo.Create(user)
}

// Delete implements [domain.IUserService].
func (u *UserService) Delete(id int) error {
	panic("unimplemented")
}

// Get implements [domain.IUserService].
func (u *UserService) Get(id int) (model.User, error) {
	panic("unimplemented")
}

// List implements [domain.IUserService].
func (s *UserService) List(page int, limit int) ([]model.User, error) {
	return s.repo.List(page, limit)
}

// UpdateProfile implements [domain.IUserService].
func (s *UserService) UpdateProfile(profile *model.Profile, userId int) error {
	return s.repo.UpdateProfile(profile, userId)
}

// ChangePassword implements [domain.IUserService].
func (s *UserService) ChangePassword(auth *model.Auth, userId int) error {
	return s.repo.ChangePassword(auth, userId)
}

func NewUserService(repo domain.IUserRepository) domain.IUserService {
	return &UserService{repo}
}
