package controllers

import (
	"net/http"

	"github.com/roronoadiogo/todo-go-rest/services/impl"
	"github.com/roronoadiogo/todo-go-rest/util/mapper"
)

func HandleGetTodoAll(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	todoImpl := &impl.TodoImpl{}
	mapper.ParseToJson(w, todoImpl.ListAllTodo())
}
