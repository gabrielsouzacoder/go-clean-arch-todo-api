package todo

import "github.com/gabrielsouzacoder/clean-new/entity"

type Reader interface {
	List() ([]*entity.Todo, error)
}

type Writer interface {
	Create(e *entity.Todo) (*entity.ID, error)
}

type Repository interface {
	Reader
	Writer
}

type UseCase interface {
	CreateTodo(description string) (*entity.ID, error)
	ListTodos() ([]*entity.Todo, error)
}