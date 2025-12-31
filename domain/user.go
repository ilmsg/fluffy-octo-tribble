package domain

import (
	"net/http"

	"github.com/ilmsg/fluffy-octo-tribble/model"
)

type IUserRepository interface {
	List(page int, limit int) ([]model.User, error)
	Create(user *model.User) error
	Get(id int) (model.User, error)
	Delete(id int) error
	Update(user *model.User) error

	UpdateProfile(profile *model.Profile, userId int) error
	ChangePassword(auth *model.Auth, userId int) error
}

type IUserService interface {
	List(page int, limit int) ([]model.User, error)
	Create(user *model.User) error
	Get(id int) (model.User, error)
	Delete(id int) error

	UpdateProfile(profile *model.Profile, userId int) error
	ChangePassword(auth *model.Auth, userId int) error
}

type IUserHandler interface {
	List(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)

	UpdateProfile(w http.ResponseWriter, r *http.Request)
	ChangePassword(w http.ResponseWriter, r *http.Request)
	ResetPassword(w http.ResponseWriter, r *http.Request)
}
