package service

import (
	"github.com/ilmsg/fluffy-octo-tribble/domain"
	"github.com/ilmsg/fluffy-octo-tribble/model"
)

type ProjectService struct {
	repo domain.IProjectRepository
}

func (s *ProjectService) Create(project *model.Project) error {
	return s.repo.Create(project)
}

func (s *ProjectService) Delete(id uint, userId uint) error {
	return s.repo.Delete(id, userId)
}

func (s *ProjectService) Get(id uint, userId uint) (model.Project, error) {
	return s.repo.Get(id, userId)
}

// func (s *ProjectService) List(page int, limit int) ([]model.Project, error) {
// 	return s.repo.List(page, limit)
// }

func (s *ProjectService) List(page int, limit int, userId uint) ([]model.Project, error) {
	return s.repo.List(page, limit, userId)
}

func (s *ProjectService) Update(project *model.Project, userId uint) error {
	return s.repo.Update(project, userId)
}

func NewProjectService(repo domain.IProjectRepository) domain.IProjectService {
	return &ProjectService{repo}
}
