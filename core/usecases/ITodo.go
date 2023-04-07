package model

import "github.com/roronoadiogo/todo-go-rest/core/model"

type ITodo interface {
	ListAllTodo() []model.Todo
}
