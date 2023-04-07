package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/roronoadiogo/todo-go-rest/controllers"
)

func RouteTodo() chi.Router {
	r := chi.NewRouter()
	r.Get("/", controllers.HandleGetTodoAll)
	return r
}
