package model

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	Tasks       []Task `json:"tasks" gorm:"foreignKey:ProjectId"`
	UserId      uint
}
