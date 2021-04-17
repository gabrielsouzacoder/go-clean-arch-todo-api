package entity

import (
	"errors"
)

type Todo struct {
	ID          ID
	Description string
	Status      bool
}

var ErrInvalidEntity = errors.New("invalid entity")

func (t Todo) Validate() interface{} {
	return nil
}

func NewTodo(description string, status bool) (*Todo, error) {
	b := &Todo{
		ID:          NewID(),
		Description: description,
		Status:      status,
	}

	err := b.Validate()

	if err != nil {
		return nil, ErrInvalidEntity
	}

	return b, nil
}
