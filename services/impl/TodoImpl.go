package impl

import (
	"github.com/roronoadiogo/todo-go-rest/core/model"
)

type TodoImpl struct {
}

func (t TodoImpl) AddTodo(model.Todo) error {
	return nil
}

func (t TodoImpl) UpdateTodo(id int) (model.Todo, error) {
	todoUpdate := model.Todo{
		ID:          1,
		Name:        "Teste todo",
		Description: "some descrition",
		Finish:      true,
		Position:    2,
	}
	return todoUpdate, nil
}

func (t TodoImpl) DeleteTodo(id int) error {
	return nil
}

func (t TodoImpl) ListAll() ([]model.Todo, error) {
	todos := []model.Todo{
		{ID: 1, Name: "Ver emails",
			Description: "Fazer as paradas",
			Finish:      false, Position: 0},
	}
	return todos, nil
}
