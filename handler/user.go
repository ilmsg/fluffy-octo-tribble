package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ilmsg/fluffy-octo-tribble/domain"
	"github.com/ilmsg/fluffy-octo-tribble/model"
	"github.com/ilmsg/fluffy-octo-tribble/util"
)

type UserWebHandler struct {
	srv domain.IUserService
}

func (h *UserWebHandler) Login(w http.ResponseWriter, r *http.Request) {
	var loginDto model.LoginDto
	if err := json.NewDecoder(r.Body).Decode(&loginDto); err != nil {
		util.GetDataResponse(w, 500, err.Error(), nil)
		return
	}

	auth, err := h.srv.Login(&loginDto)
	if err != nil {
		util.GetDataResponse(w, 500, err.Error(), nil)
		return
	}

	token, err := util.CreateToken(auth.UserId)
	if err != nil {
		util.GetDataResponse(w, 500, err.Error(), nil)
		return
	}

	loginRes := model.LoginRes{Token: token}
	json.NewEncoder(w).Encode(loginRes)
}

func (h *UserWebHandler) Register(w http.ResponseWriter, r *http.Request) {
	var registerDto model.RegisterDto
	if err := json.NewDecoder(r.Body).Decode(&registerDto); err != nil {
		util.GetDataResponse(w, 500, err.Error(), nil)
		return
	}

	user, err := h.srv.Register(&registerDto)
	if err != nil {
		util.GetDataResponse(w, 500, err.Error(), nil)
		return
	}

	token, err := util.CreateToken(user.ID)
	if err != nil {
		util.GetDataResponse(w, 500, err.Error(), nil)
		return
	}

	loginRes := model.LoginRes{Token: token}
	json.NewEncoder(w).Encode(loginRes)
}

func (h *UserWebHandler) ChangePassword(w http.ResponseWriter, r *http.Request) {
	userId, err := util.GetValueFromContext(r.Context())
	if err != nil {
		util.GetDataResponse(w, 500, err.Error(), nil)
		return
	}

	var changePasswordDto model.ChangePasswordDto
	if err := json.NewDecoder(r.Body).Decode(&changePasswordDto); err != nil {
		util.GetDataResponse(w, 500, err.Error(), nil)
		return
	}

	// validate password dto
	user, err := h.srv.Get(userId)
	if err != nil {
		util.GetDataResponse(w, 500, err.Error(), nil)
		return
	}

	auth := user.Auth
	auth.Password = changePasswordDto.NewPassword
	if err := h.srv.ChangePassword(auth, userId); err != nil {
		util.GetDataResponse(w, 500, err.Error(), nil)
		return
	}

	newPassword := model.NewPassword{
		Password: changePasswordDto.NewPassword,
	}
	json.NewEncoder(w).Encode(newPassword)
}

func (h *UserWebHandler) ResetPassword(w http.ResponseWriter, r *http.Request) {
	var resetPasswordDto model.ResetPasswordDto
	if err := json.NewDecoder(r.Body).Decode(&resetPasswordDto); err != nil {
		util.GetDataResponse(w, 500, err.Error(), nil)
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
		util.GetDataResponse(w, 500, err.Error(), nil)
		return
	}

	// validate RegisterDto

	newUser := model.User{
		Auth: &model.Auth{
			Email:    registerDto.Email,
			Password: registerDto.Password,
		},
		Profile: &model.Profile{
			Name:     registerDto.Name,
			Location: registerDto.Location,
		},
	}
	if err := h.srv.Create(&newUser); err != nil {
		util.GetDataResponse(w, 500, err.Error(), nil)
		return
	}

	json.NewEncoder(w).Encode(newUser)
}

// Delete implements [domain.IUserHandler].
func (h *UserWebHandler) Delete(w http.ResponseWriter, r *http.Request) {

}

func (h *UserWebHandler) Get(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("user_id").(uint)

	// id, _ := strconv.Atoi(mux.Vars(r)["id"])
	user, err := h.srv.Get(userId)
	if err != nil {
		util.GetDataResponse(w, 500, err.Error(), nil)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func (h *UserWebHandler) List(w http.ResponseWriter, r *http.Request) {
	page, limit := util.PageLimit(r)
	users, err := h.srv.List(page, limit)
	if err != nil {
		util.GetDataResponse(w, 500, err.Error(), nil)
		return
	}

	json.NewEncoder(w).Encode(users)
}

func (h *UserWebHandler) GetProfile(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("user_id").(uint)
	log.Printf("UserId: %d", userId)

	profile, err := h.srv.GetProfile(userId)
	if err != nil {
		util.GetDataResponse(w, 500, err.Error(), nil)
		return
	}

	util.GetDataResponse(w, 200, "Profile retrieved successfully", profile)
}

func (h *UserWebHandler) UpdateProfile(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("user_id").(uint)

	var profileDto model.ProfileDto
	if err := json.NewDecoder(r.Body).Decode(&profileDto); err != nil {
		util.GetDataResponse(w, 500, err.Error(), nil)
		return
	}

	newProfile := model.Profile{Name: profileDto.Name, Location: profileDto.Location}
	err := h.srv.UpdateProfile(&newProfile, userId)
	if err != nil {
		util.GetDataResponse(w, 500, err.Error(), nil)
		return
	}

	profile, err := h.srv.GetProfile(userId)
	if err != nil {
		util.GetDataResponse(w, 500, err.Error(), nil)
		return
	}
	util.GetDataResponse(w, 200, "Profile updated successfully", profile)
}

func NewUserWebHandler(srv domain.IUserService) domain.IUserHandler {
	return &UserWebHandler{srv}
}
