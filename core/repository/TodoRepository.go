package repository

import "github.com/roronoadiogo/todo-go-rest/core/model"

type TodoRepository interface {
	ListAll() []model.Todo
}
