package impl

import (
	"github.com/roronoadiogo/todo-go-rest/core/model"
)

type TodoImpl struct {
}

func (t TodoImpl) ListAll() []model.Todo {
	todos := []model.Todo{
		{ID: 1, Name: "Ver emails",
			Description: "Fazer as paradas",
			Finish:      false, Position: 0},
	}
	return todos
}
