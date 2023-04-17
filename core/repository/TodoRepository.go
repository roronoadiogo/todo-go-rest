package repository

import "github.com/roronoadiogo/todo-go-rest/core/model"

type TodoRepository interface {
	AddTodo(model.Todo) error
	UpdateTodo(id int) (model.Todo, error)
	DeleteTodo(id int) error
	ListAll() ([]model.Todo, error)
}
