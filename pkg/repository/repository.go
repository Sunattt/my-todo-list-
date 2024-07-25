package repository

type Authorization interface {
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

func NewRepository() *Repository {
	return &Repository{}
}
