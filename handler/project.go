package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/ilmsg/fluffy-octo-tribble/domain"
	"github.com/ilmsg/fluffy-octo-tribble/model"
	"github.com/ilmsg/fluffy-octo-tribble/util"
	"gorm.io/gorm"
)

type ProjectHandler struct {
	srv domain.IProjectService
}

func (h *ProjectHandler) Create(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("user_id").(uint)

	var projectDto model.ProjectDto
	if err := json.NewDecoder(r.Body).Decode(&projectDto); err != nil {
		util.GetDataResponse(w, 500, err.Error(), nil)
		return
	}

	// validate project dto

	newProject := model.Project{
		Title:       projectDto.Title,
		Description: projectDto.Description,
		Tasks:       []model.Task{},
		UserId:      userId,
	}
	if err := h.srv.Create(&newProject); err != nil {
		util.GetDataResponse(w, 500, err.Error(), nil)
		return
	}

	util.GetDataResponse(w, 201, "Project created successfully", newProject)
}

func (h *ProjectHandler) Delete(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("user_id").(uint)
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	err := h.srv.Delete(uint(id), userId)
	if err != nil {
	}

	util.GetDataResponse(w, 200, "Project delete successfully", nil)
}

func (h *ProjectHandler) Get(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("user_id").(uint)
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	project, err := h.srv.Get(uint(id), userId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			util.GetDataResponse(w, 404, "Data not found", nil)
			return
		}

		util.GetDataResponse(w, 400, err.Error(), nil)
		return
	}

	util.GetDataResponse(w, 200, "Data found", project)
}

// func (h *ProjectHandler) List(w http.ResponseWriter, r *http.Request) {
// 	// userId := r.Context().Value("user_id").(uint)
// 	userId := r.Context().Value("user_id").(uint)
// 	page, limit := util.PageLimit(r)

// 	projects, err := h.srv.List(page, limit, userId)
// 	if err != nil {
// 		util.GetDataResponse(w, 500, err.Error(), nil)
// 		return
// 	}

// 	dataRes := model.DataResponse{
// 		Status:  200,
// 		Message: "Projects retrieved successfully",
// 		Data:    projects,
// 	}
// 	json.NewEncoder(w).Encode(dataRes)
// }

func (h *ProjectHandler) List(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("user_id").(uint)
	page, limit := util.PageLimit(r)

	projects, err := h.srv.List(page, limit, userId)
	if err != nil {
		util.GetDataResponse(w, 500, err.Error(), nil)
		return
	}

	util.GetDataResponse(w, 200, "Projects retrieved successfully", projects)
}

func (h *ProjectHandler) Update(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("user_id").(uint)
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	var projectDto model.ProjectDto
	if err := json.NewDecoder(r.Body).Decode(&projectDto); err != nil {
		util.GetDataResponse(w, 500, err.Error(), nil)
		return
	}

	updateProject := model.Project{
		ID:          uint(id),
		Title:       projectDto.Title,
		Description: projectDto.Description,
	}
	if err := h.srv.Update(&updateProject, userId); err != nil {
		util.GetDataResponse(w, 500, err.Error(), nil)
		return
	}

	project, err := h.srv.Get(uint(id), userId)
	if err != nil {
		util.GetDataResponse(w, 500, err.Error(), nil)
		return
	}
	util.GetDataResponse(w, 200, "Project update successfully", project)
}

func NewProjectHandler(srv domain.IProjectService) domain.IProjectHandler {
	return &ProjectHandler{srv}
}
