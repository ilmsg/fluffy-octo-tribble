package domain

import (
	"net/http"

	"github.com/ilmsg/fluffy-octo-tribble/model"
)

type IUserRepository interface {
	List(page int, limit int) ([]model.User, error)
	Get(id uint) (model.User, error)
	Delete(id uint) error
	Update(user *model.User) error

	// Exists(email string) (isExists bool)

	Create(user *model.User) error
	Register(registerDto *model.RegisterDto) (model.User, error)
	Login(loginDto *model.LoginDto) (model.Auth, error)

	GetProfile(id uint) (model.Profile, error)
	UpdateProfile(profile *model.Profile, userId uint) error
	ChangePassword(auth *model.Auth, userId uint) error
}

type IUserService interface {
	List(page int, limit int) ([]model.User, error)
	Get(id uint) (model.User, error)
	Delete(id uint) error

	Create(user *model.User) error
	Register(registerDto *model.RegisterDto) (model.User, error)
	Login(loginDto *model.LoginDto) (model.Auth, error)

	GetProfile(id uint) (model.Profile, error)
	UpdateProfile(profile *model.Profile, userId uint) error
	ChangePassword(auth *model.Auth, userId uint) error
}

type IUserHandler interface {
	List(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)

	Create(w http.ResponseWriter, r *http.Request)
	Register(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)

	GetProfile(w http.ResponseWriter, r *http.Request)
	UpdateProfile(w http.ResponseWriter, r *http.Request)
	ChangePassword(w http.ResponseWriter, r *http.Request)
	ResetPassword(w http.ResponseWriter, r *http.Request)
}
