package todo

import (
	"errors"
	"fmt"

	"github.com/Marrquito/GoToDoAPP/database/repositories"
)

type DeleteToDoService interface {
	Execute(toDoOptions DeleteToDoServiceOptions) error
}

type deleteToDoService struct {
	toDoRepository repositories.ToDoRepository
}

type DeleteToDoServiceOptions struct {
	ID        uint
}

func NewDeleteToDoService() DeleteToDoService {
	return &deleteToDoService{
		toDoRepository: repositories.NewToDoRepository(),
	}
}

func (s deleteToDoService) Execute(options DeleteToDoServiceOptions) error {
	if _, readError := s.toDoRepository.Read(options.ID); readError != nil {
		errMsg := fmt.Sprintf("Ocorreu um erro ao buscar ToDo, verifique o ID e tente novamente. %s", readError.Error())
		return errors.New(errMsg)
	}
	
	if deleteError := s.toDoRepository.Delete(options.ID); deleteError != nil {
		errMsg := fmt.Sprintf("Ocorreu um erro ao deletar ToDo, tente novamente. %s", deleteError.Error())
		return errors.New(errMsg)
	}

	return nil
}
