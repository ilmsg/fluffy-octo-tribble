package model

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Title      string     `json:"title"`
	TaskStatus TaskStatus `json:"task_status"`
	UserId     uint       `json:"user_id"`
	ProjectId  uint       `json:"project_id"`
}
