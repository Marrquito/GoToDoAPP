package todo

import (
	"errors"
	"fmt"

	"github.com/Marrquito/GoToDoAPP/api/requests"
	"github.com/Marrquito/GoToDoAPP/database/models"
	"github.com/Marrquito/GoToDoAPP/database/repositories"
)

type CreateToDoService interface {
	Execute(toDo CreateToDoServiceOptions) (models.ToDo, error)
}

type createToDoService struct {
	toDoRepository repositories.ToDoRepository
}

type CreateToDoServiceOptions struct {
	ToDo       *requests.CreateToDoRequest
}

func NewCreateToDoService() CreateToDoService {
	return &createToDoService{
		toDoRepository: repositories.NewToDoRepository(),
	}
}

func (s createToDoService) Execute(options CreateToDoServiceOptions) (models.ToDo, error) {
	if options.ToDo == nil {
		return models.ToDo{}, errors.New("ToDo == nil")
	}
	
	toDo, err := options.ToDo.ToModel()
	if err != nil {
		errMsg := fmt.Sprintf("Ocorreu um erro ao criar ToDo, verifique os dados e tente novamente. %s", err.Error())
		return models.ToDo{}, errors.New(errMsg)
	}
	
	createdToDo, createError := s.toDoRepository.Create(toDo)
	if createError != nil {
		errMsg := fmt.Sprintf("Ocorreu um erro ao criar ToDo, tente novamente. %s", createError.Error())
		return models.ToDo{}, errors.New(errMsg)
	}

	return createdToDo, nil
}