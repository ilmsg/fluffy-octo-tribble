package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/ilmsg/fluffy-octo-tribble/domain"
	"github.com/ilmsg/fluffy-octo-tribble/model"
	"github.com/ilmsg/fluffy-octo-tribble/util"
)

type UserWebHandler struct {
	srv domain.IUserService
}

// ChangePassword implements [domain.IUserHandler].
func (h *UserWebHandler) ChangePassword(w http.ResponseWriter, r *http.Request) {
	userId, err := util.GetValueFromContext(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var changePasswordDto model.ChangePasswordDto
	if err := json.NewDecoder(r.Body).Decode(&changePasswordDto); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// validate password dto
	user, err := h.srv.Get(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	auth := &user.Auth
	auth.Password = changePasswordDto.NewPassword
	if err := h.srv.ChangePassword(auth, userId); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	newPassword := model.NewPassword{
		Password: changePasswordDto.NewPassword,
	}
	json.NewEncoder(w).Encode(newPassword)
}

// ResetPassword implements [domain.IUserHandler].
func (h *UserWebHandler) ResetPassword(w http.ResponseWriter, r *http.Request) {
	var resetPasswordDto model.ResetPasswordDto
	if err := json.NewDecoder(r.Body).Decode(&resetPasswordDto); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// validate email

	// token reset password
	tokenResetPassword := struct{ Token string }{Token: "xxxx"}
	json.NewEncoder(w).Encode(tokenResetPassword)
}

// Create implements [domain.IUserHandler].
func (h *UserWebHandler) Create(w http.ResponseWriter, r *http.Request) {
	var registerDto model.RegisterDto
	if err := json.NewDecoder(r.Body).Decode(&registerDto); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// validate RegisterDto

	newUser := model.User{
		Auth: model.Auth{
			Email:    registerDto.Email,
			Password: registerDto.Password,
		},
		Profile: model.Profile{
			Name:     registerDto.Name,
			Location: registerDto.Location,
		},
	}
	if err := h.srv.Create(&newUser); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(newUser)
}

// Delete implements [domain.IUserHandler].
func (h *UserWebHandler) Delete(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

// Get implements [domain.IUserHandler].
func (h *UserWebHandler) Get(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	user, err := h.srv.Get(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)
}

// List implements [domain.IUserHandler].
func (h *UserWebHandler) List(w http.ResponseWriter, r *http.Request) {
	page, limit := util.PageLimit(r)
	users, err := h.srv.List(page, limit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(users)
}

// Update implements [domain.IUserHandler].
func (h *UserWebHandler) UpdateProfile(w http.ResponseWriter, r *http.Request) {
	var profileDto model.ProfileDto
	if err := json.NewDecoder(r.Body).Decode(&profileDto); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	user, err := h.srv.Get(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func NewUserWebHandler(srv domain.IUserService) domain.IUserHandler {
	return &UserWebHandler{srv}
}
