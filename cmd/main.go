package main

import (
	"log"
	todo_list "todo"
	"todo/pkg/handler"
	"todo/pkg/repository"
	"todo/pkg/service"
)

func main() {

	repos := repository.NewRepository()
	services := service.NewService(*repos)
	handlers := handler.NewHandler(services)

	srv := new(todo_list.Server)
	if err := srv.Run("8000", handlers.InitRoute()); err != nil {
		log.Fatal("")
	}
}
