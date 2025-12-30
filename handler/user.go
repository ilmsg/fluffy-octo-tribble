package handler

import (
	"net/http"

	"github.com/ilmsg/fluffy-octo-tribble/domain"
)

type UserWebHandler struct {
	srv domain.IUserService
}

// Create implements [domain.IUserHandler].
func (u *UserWebHandler) Create(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

// Delete implements [domain.IUserHandler].
func (u *UserWebHandler) Delete(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

// Get implements [domain.IUserHandler].
func (u *UserWebHandler) Get(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

// List implements [domain.IUserHandler].
func (u *UserWebHandler) List(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

// Update implements [domain.IUserHandler].
func (u *UserWebHandler) Update(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

func NewUserWebHandler(srv domain.IUserService) domain.IUserHandler {
	return &UserWebHandler{srv}
}
