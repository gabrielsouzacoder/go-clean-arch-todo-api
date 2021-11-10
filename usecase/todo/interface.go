package todo

import "github.com/gabrielsouzacoder/clean-new/entity"

type Reader interface {
	List() ([]*entity.Todo, error)
	FindById(id *entity.ID) *entity.Todo
}

type Writer interface {
	Create(e *entity.Todo) (*entity.ID, error)
	Delete(id *entity.ID) error
}

type Repository interface {
	Reader
	Writer
	Update(todo *entity.Todo) *entity.Todo
}

type UseCase interface {
	CreateTodo(description string) (*entity.ID, error)
	ListTodos() ([]*entity.Todo, error)
}
