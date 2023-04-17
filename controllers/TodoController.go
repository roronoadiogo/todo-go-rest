package controllers

import (
	"net/http"

	"github.com/roronoadiogo/todo-go-rest/core/usecases"
	"github.com/roronoadiogo/todo-go-rest/services/impl"
	"github.com/roronoadiogo/todo-go-rest/util/mapper"
)

func HandleGetTodoAll(w http.ResponseWriter, r *http.Request) {
	u := usecases.NewUseCaseTodo(impl.TodoImpl{})

	todos, err := u.Execute()
	if err != nil {
		http.Error(w, "Error getting todos in repository", http.StatusInternalServerError)
		return
	}

	mapper.ParseToJson(w, todos)
}
