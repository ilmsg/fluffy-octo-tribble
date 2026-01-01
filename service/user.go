package service

import (
	"github.com/ilmsg/fluffy-octo-tribble/domain"
	"github.com/ilmsg/fluffy-octo-tribble/model"
)

type UserService struct {
	repo domain.IUserRepository
}

// Login implements [domain.IUserService].
func (s *UserService) Login(loginDto *model.LoginDto) (model.Auth, error) {
	return s.repo.Login(loginDto)
}

// Register implements [domain.IUserService].
func (s *UserService) Register(registerDto *model.RegisterDto) (model.User, error) {
	return s.repo.Register(registerDto)
}

// Create implements [domain.IUserService].
func (s *UserService) Create(user *model.User) error {
	return s.repo.Create(user)
}

// Delete implements [domain.IUserService].
func (s *UserService) Delete(id uint) error {
	return s.repo.Delete(id)
}

// Get implements [domain.IUserService].
func (s *UserService) Get(id uint) (model.User, error) {
	return s.repo.Get(id)
}

// List implements [domain.IUserService].
func (s *UserService) List(page int, limit int) ([]model.User, error) {
	return s.repo.List(page, limit)
}

// GetProfile implements [domain.IUserService].
func (s *UserService) GetProfile(id uint) (model.Profile, error) {
	return s.repo.GetProfile(id)
}

// UpdateProfile implements [domain.IUserService].
func (s *UserService) UpdateProfile(profile *model.Profile, userId uint) error {
	return s.repo.UpdateProfile(profile, userId)
}

// ChangePassword implements [domain.IUserService].
func (s *UserService) ChangePassword(auth *model.Auth, userId uint) error {
	return s.repo.ChangePassword(auth, userId)
}

func NewUserService(repo domain.IUserRepository) domain.IUserService {
	return &UserService{repo}
}
