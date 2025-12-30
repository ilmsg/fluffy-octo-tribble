package repository

import (
	"github.com/ilmsg/fluffy-octo-tribble/domain"
	"github.com/ilmsg/fluffy-octo-tribble/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
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

func (repo *UserRepository) Update(user *model.User) error {
	return repo.db.Save(user).Error
}

func (repo *UserRepository) Delete(id int) error {
	return repo.db.Delete(&model.User{}, id).Error
}

func NewUserRepository(db *gorm.DB) domain.IUserRepository {
	return &UserRepository{db}
}
