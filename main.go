package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/roronoadiogo/todo-go-rest/config"
	"github.com/roronoadiogo/todo-go-rest/routes"
	"go.uber.org/zap"
)

var logger *zap.Logger = config.InitializeLogger()

func init() {
	logger.Info("Started Application Todo Go List")
	config.LoadEnv(logger)
}

func main() {

	server := chi.NewRouter()
	server.Use(middleware.Logger)

	server.Mount("/todo", routes.RouteTodo())
	http.ListenAndServe(":3000", server)

}
