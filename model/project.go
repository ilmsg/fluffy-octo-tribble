package model

type Project struct {
	// gorm.Model
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Tasks       []Task `json:"tasks" gorm:"foreignKey:ProjectId"`
	UserId      uint   `json:"user_id"`
}

type ProjectDto struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type ProjectRes struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Tasks       []Task `json:"tasks" gorm:"foreignKey:ProjectId"`
	UserId      uint   `json:"user_id"`
}

type DataResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}
