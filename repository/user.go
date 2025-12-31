package repository

import (
	"fmt"
	"log"

	"github.com/ilmsg/fluffy-octo-tribble/domain"
	"github.com/ilmsg/fluffy-octo-tribble/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

// func (repo *UserRepository) Exists(email string) (isExists bool) {
// 	var auth model.Auth
// 	err := repo.db.First(&auth, "email = ?", email).Error
// 	if err != nil {
// 		return err.Error() != "record not found"
// 	}
// 	return true
// }

func (repo *UserRepository) Login(loginDto *model.LoginDto) (model.Auth, error) {
	var auth model.Auth
	tx := repo.db.Where("email = ?", loginDto.Email).First(&auth)
	if tx.Error == gorm.ErrRecordNotFound {
		return auth, fmt.Errorf("user not found.")
	}
	return auth, tx.Error
}

func (repo *UserRepository) Register(registerDto *model.RegisterDto) (model.User, error) {
	tx := repo.db.Where("Email = ?", registerDto.Email).First(&model.Auth{})
	if tx.Error == nil {
		return model.User{}, fmt.Errorf("user exists")
	} else if tx.Error != gorm.ErrRecordNotFound {
		return model.User{}, tx.Error
	}

	user := model.User{}
	tx = repo.db.Create(&user)

	auth := model.Auth{Email: registerDto.Email, Password: registerDto.Password, UserId: user.ID}
	if tx := repo.db.Create(&auth); tx.Error != nil {
		log.Fatal(tx.Error)
		return model.User{}, tx.Error
	}

	profile := model.Profile{Name: registerDto.Name, Location: registerDto.Location, UserId: user.ID}
	if tx := repo.db.Create(&profile); tx.Error != nil {
		return model.User{}, tx.Error
	}

	var newUser model.User
	tx = repo.db.Preload("Auth").Preload("Profile").Where("ID = ?", user.ID).First(&newUser)
	return newUser, tx.Error
}

func (repo *UserRepository) List(page int, limit int) ([]model.User, error) {
	var users []model.User
	tx := repo.db.Offset((page - 1) * limit).Limit(limit).Find(&users)
	return users, tx.Error
}

func (repo *UserRepository) Create(user *model.User) error {
	return repo.db.Create(&user).Error
}

func (repo *UserRepository) Get(id int) (model.User, error) {
	var user model.User
	tx := repo.db.First(&user, id)
	return user, tx.Error
}

func (repo *UserRepository) Delete(id int) error {
	return repo.db.Delete(&model.User{}, id).Error
}

func (repo *UserRepository) Update(user *model.User) error {
	return repo.db.Save(user).Error
}

func (repo *UserRepository) UpdateProfile(profile *model.Profile, userId int) error {
	return repo.db.Where("userId = ?", userId).Updates(profile).Error
}

func (repo *UserRepository) ChangePassword(auth *model.Auth, userId int) error {
	return repo.db.Where("userId = ?", userId).Updates(auth).Error
}

func NewUserRepository(db *gorm.DB) domain.IUserRepository {
	return &UserRepository{db}
}
