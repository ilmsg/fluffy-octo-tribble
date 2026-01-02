package repository

import (
	"github.com/ilmsg/fluffy-octo-tribble/domain"
	"github.com/ilmsg/fluffy-octo-tribble/model"
	"gorm.io/gorm"
)

type ProjectRepository struct {
	db *gorm.DB
}

func (r *ProjectRepository) Create(project *model.Project) error {
	return r.db.Create(&project).Error
}

func (r *ProjectRepository) Delete(id uint, userId uint) error {
	return r.db.Where("user_id = ?", userId).Delete(&model.Project{}, "id = ?", id).Error
}

func (r *ProjectRepository) Get(id uint, userId uint) (model.Project, error) {
	var project model.Project
	tx := r.db.Preload("Tasks").Where("user_id = ?", userId).First(&project, "id = ?", id)
	return project, tx.Error
}

// func (r *ProjectRepository) List(page int, limit int) ([]model.Project, error) {
// 	var projects []model.Project
// 	tx := r.db.Offset((page - 1) * limit).Limit(limit).Find(&projects)
// 	return projects, tx.Error
// }

func (r *ProjectRepository) List(page int, limit int, user_id uint) ([]model.Project, error) {
	var projects []model.Project
	tx := r.db.Preload("Tasks").Offset((page-1)*limit).Limit(limit).Find(&projects, "user_id = ?", user_id)
	return projects, tx.Error
}

func (r *ProjectRepository) Update(project *model.Project, userId uint) error {
	return r.db.Where("id = ?", project.ID).Where("user_id = ?", userId).Updates(&project).Error
}

func NewProjectRepository(db *gorm.DB) domain.IProjectRepository {
	return &ProjectRepository{db}
}
