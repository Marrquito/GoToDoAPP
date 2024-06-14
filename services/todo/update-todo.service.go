package todo

import (
	"errors"
	"fmt"

	"github.com/Marrquito/GoToDoAPP/api/requests"
	"github.com/Marrquito/GoToDoAPP/database/models"
	"github.com/Marrquito/GoToDoAPP/database/repositories"
)

type UpdateToDoService interface {
	Execute(options UpdateToDoServiceOptions) (*models.ToDo, error)
}

type updateToDoService struct {
	toDoRepository repositories.ToDoRepository
}

type UpdateToDoServiceOptions struct {
	ID        uint
	Data      *requests.UpdateToDoRequest
}

func NewUpdateToDoService() UpdateToDoService {
	return &updateToDoService{
		toDoRepository: repositories.NewToDoRepository(),
	}
}

func (s updateToDoService) Execute(options UpdateToDoServiceOptions) (*models.ToDo, error) {
	toDo, readError := s.toDoRepository.Read(options.ID)
	if readError != nil {
		errMsg := fmt.Sprintf("Ocorreu um erro ao atualizar encontrar, tente novamente. %s", readError.Error())
		return nil, errors.New(errMsg)
	}

	options.Data.Validate(toDo)
	_, updateError := s.toDoRepository.Update(*toDo)
	if updateError != nil {
		errMsg := fmt.Sprintf("Ocorreu um erro ao atualizar ToDo, tente novamente. %s", updateError.Error())
		return nil, errors.New(errMsg)
	}

	return toDo, nil
}