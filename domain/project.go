package domain

import (
	"net/http"

	"github.com/ilmsg/fluffy-octo-tribble/model"
)

type IProjectRepository interface {
	Create(project *model.Project) error
	// List(page int, limit int) ([]model.Project, error)
	List(page int, limit int, userId uint) ([]model.Project, error)
	Get(id uint, userId uint) (model.Project, error)
	Delete(id uint, userId uint) error
	Update(project *model.Project, userId uint) error
}

type IProjectService interface {
	Create(project *model.Project) error
	// List(page int, limit int) ([]model.Project, error)
	List(page int, limit int, userId uint) ([]model.Project, error)
	Get(id uint, userId uint) (model.Project, error)
	Delete(id uint, userId uint) error
	Update(project *model.Project, userId uint) error
}

type IProjectHandler interface {
	Create(w http.ResponseWriter, r *http.Request)
	List(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
}
