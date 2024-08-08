package service

import (
	repository "todo/internal/repository/psgdb"
	"todo/pkg/todo_list"
)

type Authorization interface {
	CreateUser(user todo_list.User) (int, error)
}  

type TodoList interface {
}

type TodoItems interface {
}

type Service struct {
	Authorization
	TodoList
	TodoItems
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
	}
}
