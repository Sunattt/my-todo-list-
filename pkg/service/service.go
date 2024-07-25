package service

import (
	"todo/pkg/repository"
)

type Authorization interface {
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

func NewService(repo repository.Repository) *Service {
	return &Service{}
}
