package main

import (
	"log"
	todo_list "todo"
	"todo/pkg/handler"
)

func main() {

	handlers := new(handler.Handler)
	srv := new(todo_list.Server)
	if err := srv.Run("8000", handlers.InitRoute()); err != nil {
		log.Fatal("")
	}
}
