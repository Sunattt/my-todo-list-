package psgdb

import (
	"todo/pkg/todo_list"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user todo_list.User) (int, error)
}

type TodoList interface {
}

type TodoItems interface {
}

type Repository struct {
	Authorization
	TodoItems
	TodoList
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
