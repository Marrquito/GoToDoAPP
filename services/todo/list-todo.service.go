package todo

import (
	"errors"
	"fmt"

	"github.com/Marrquito/GoToDoAPP/api/requests"
	"github.com/Marrquito/GoToDoAPP/database/models"
	"github.com/Marrquito/GoToDoAPP/database/repositories"
)

type ListToDoService interface {
	Execute(options *ListToDoServiceOptions) ([]models.ToDo, *int64, error)
}

type listToDoService struct {
	toDoRepository repositories.ToDoRepository
}

type ListToDoServiceOptions struct {
	Filters   *requests.ListToDoRequest
}

func NewListToDoService() ListToDoService {
	return &listToDoService{
		toDoRepository: repositories.NewToDoRepository(),
	}
}

func (s listToDoService) Execute(options *ListToDoServiceOptions) ([]models.ToDo, *int64, error) {
	page := 1
	limit := -1
	order := "desc"

	if options != nil {

		if options.Filters.Limit > 0 {
			limit = options.Filters.Limit
		}

		if options.Filters.Order != "" {
			order = options.Filters.Order
		}

		if options.Filters.Page > 0 {
			page = options.Filters.Page
		}
	}

	toDos, total, err := s.toDoRepository.List(&repositories.ListToDoOptions{
		Page:  page,
		Limit: limit,
		Order: order,
	})
	if err != nil {
		errMsg := fmt.Sprintf("Ocorreu um erro ao Listar ToDos, tente novamente. %s", err.Error())
		return []models.ToDo{}, nil, errors.New(errMsg)
	}

	return toDos, &total, nil
}
