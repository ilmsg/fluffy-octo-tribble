package model

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Auth    *Auth
	Profile *Profile
}

type Auth struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	UserId   uint   `json:"user_id"`
}

func (auth *Auth) BeforeCreate(tx *gorm.DB) (err error) {
	err = tx.First(&Auth{}, "email = ?", auth.Email).Error
	if (err != nil && err.Error() != "record not found") || err == nil {
		err = fmt.Errorf("user exists")
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(auth.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	auth.Password = string(hash)
	return nil
}

func (auth *Auth) BeforeUpdate(tx *gorm.DB) (err error) {
	if auth.Password != "" {
		hash, err := bcrypt.GenerateFromPassword([]byte(auth.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		tx.UpdateColumn("Password", hash)
	}
	return
}

type Profile struct {
	Name     string `json:"name"`
	Location string `json:"location"`
	UserId   uint   `json:"user_id"`
}

type RegisterDto struct {
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`

	Name     string `json:"name"`
	Location string `json:"location"`
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

type NewPassword struct {
	Password string
}

type ResetPasswordDto struct {
	Email string
}

type Token struct {
	AccessToken string
}

type LoginRes struct {
	Token string
}
