package requests

import "github.com/Marrquito/GoToDoAPP/database/models"

type CreateToDoRequest struct {
	Title       string `json:"title" form:"title"`
	Description string `json:"description" form:"description"`
}

type ListToDoRequest struct {
	Page  int    `form:"page"`
	Limit int    `form:"limit"`
	Order string `form:"order"`
}

type UpdateToDoRequest struct {
	Title       *string `json:"title" form:"title"`
	Description *string `json:"description" form:"description"`
	Done        *bool   `json:"done" form:"done"`
}

func (ToDo *CreateToDoRequest) ToModel() (models.ToDo, error) {
	return models.ToDo{
		Title:       ToDo.Title,
		Description: ToDo.Description,
	}, nil
}

func (ToDo *UpdateToDoRequest) ToModel() (models.ToDo, error) {
	return models.ToDo{
		Title:       *ToDo.Title,
		Description: *ToDo.Description,
		Done:        *ToDo.Done,
	}, nil
}

func (ToDo *UpdateToDoRequest) Validate(model *models.ToDo) {
	if ToDo.Title != nil {
		model.Title = *ToDo.Title
	}
	if ToDo.Description != nil {
		model.Description = *ToDo.Description
	}
	if ToDo.Done != nil {
		model.Done = *ToDo.Done
	}
}
