package controllers

import (
	"net/http"

	"github.com/roronoadiogo/todo-go-rest/core/usecases"
	"github.com/roronoadiogo/todo-go-rest/services/impl"
	"github.com/roronoadiogo/todo-go-rest/util/mapper"
)

func HandleGetTodoAll(w http.ResponseWriter, r *http.Request) {
	u := usecases.NewUseCaseTodo(impl.TodoImpl{})
	mapper.ParseToJson(w, u.Execute())
}
