package todo

import (
	"github.com/gabrielsouzacoder/clean-new/entity"
)

type Service struct {
	repo Repository
}

func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

func (s *Service) CreateTodo(description string) (*entity.ID, error) {
	b, err := entity.NewTodo(description, false)

	if err != nil {
		return nil, err
	}
	return s.repo.Create(b)
}

func (s *Service) ListTodos() ([]*entity.Todo, error) {
	todos, _ := s.repo.List()

	return todos, nil
}
