package todo

import (
	"errors"
	"fmt"

	"github.com/Marrquito/GoToDoAPP/database/models"
	"github.com/Marrquito/GoToDoAPP/database/repositories"
)

type ReadToDoService interface {
	Execute(readToDoOptions ReadToDoServiceOptions) (*models.ToDo, error)
}

type readToDoService struct {
	toDoRepository repositories.ToDoRepository
}

type ReadToDoServiceOptions struct {
	ID uint
}

func NewReadToDoService() ReadToDoService {
	return &readToDoService{
		toDoRepository: repositories.NewToDoRepository(),
	}
}

func (s readToDoService) Execute(options ReadToDoServiceOptions) (*models.ToDo, error) {
	toDo, err := s.toDoRepository.Read(options.ID)
	if err != nil {
		errMsg := fmt.Sprintf("Ocorreu um erro ao ler ToDo, tente novamente. %s", err.Error())
		return nil, errors.New(errMsg)
	}

	return toDo, nil
}