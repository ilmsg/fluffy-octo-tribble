package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Auth    Auth
	Profile Profile
}

type Auth struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Profile struct {
	Name     string `json:"name"`
	Location string `json:"location"`
	UserId   uint
}

type RegisterDto struct {
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

type LoginDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ProfileDto struct {
	Name     string `json:"name"`
	Location string `json:"location"`
}

type ChangePasswordDto struct {
	NewPassword        string
	ConfirmNewPassword string
}

type PasswordRecoveryDto struct {
	Email string
}

type Token struct {
	AccessToken string
}
