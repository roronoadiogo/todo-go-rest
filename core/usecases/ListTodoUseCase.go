package usecases

import (
	"github.com/roronoadiogo/todo-go-rest/core/model"
	"github.com/roronoadiogo/todo-go-rest/core/repository"
)

type ListTodoUseCase struct {
	repo repository.TodoRepository
}

func NewUseCaseTodo(repo repository.TodoRepository) ListTodoUseCase {
	return ListTodoUseCase{repo: repo}
}

func (u *ListTodoUseCase) Execute() ([]model.Todo, error) {

	todos, err := u.repo.ListAll()
	if err != nil {
		return nil, err
	}
	return todos, nil
}
