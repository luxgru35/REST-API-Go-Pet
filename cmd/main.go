package main

import (
	"log"

	"github.com/luxgru35/todo-app"
	"github.com/luxgru35/todo-app/pkg/handler"
	"github.com/luxgru35/todo-app/pkg/repository"
	"github.com/luxgru35/todo-app/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(todo.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("failed to run server: %s", err.Error())
	}
}
