package responses

import (
	"time"

	"github.com/Marrquito/GoToDoAPP/database/models"
)

type CreateToDoResponse struct {
	ID          uint   `json:"id" form:"id"`
	Title       string `json:"title" form:"title"`
	Description string `json:"description" form:"description"`
	Done        bool   `json:"done" form:"done"`
}

type ReadToDoResponse struct {
	ID          uint      `json:"id" form:"id"`
	Title       string    `json:"title" form:"title"`
	Description string    `json:"description" form:"description"`
	Done        bool      `json:"done" form:"done"`
	CreatedAt   time.Time `json:"created_at" form:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" form:"updated_at"`
}

type ListToDoResponse struct {
	ID          uint      `json:"id" form:"id"`
	Title       string    `json:"title" form:"title"`
	Description string    `json:"description" form:"description"`
	Done        bool      `json:"done" form:"done"`
	CreatedAt   time.Time `json:"created_at" form:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" form:"updated_at"`
}

func NewCreateToDoResponse(model models.ToDo) CreateToDoResponse {
	return CreateToDoResponse{
		ID:          model.ID,
		Title:       model.Title,
		Description: model.Description,
		Done:        model.Done,
	}
}

func NewReadToDoResponse(model models.ToDo) ReadToDoResponse {
	return ReadToDoResponse{
		ID:          model.ID,
		Title:       model.Title,
		Description: model.Description,
		Done:        model.Done,
		CreatedAt:   model.CreatedAt,
		UpdatedAt:   model.UpdatedAt,
	}
}

func NewListToDoResponse(model models.ToDo) ListToDoResponse {
	return ListToDoResponse{
		ID:          model.ID,
		Title:       model.Title,
		Description: model.Description,
		Done:        model.Done,
		CreatedAt:   model.CreatedAt,
		UpdatedAt:   model.UpdatedAt,
	}
}

func NewListToDoResponses(models []models.ToDo) []ListToDoResponse {
	ToDoResponses := []ListToDoResponse{}

	for _, model := range models {
		ToDoResponses = append(ToDoResponses, NewListToDoResponse(model))
	}

	return ToDoResponses
}
