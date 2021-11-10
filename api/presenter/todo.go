package presenter

import "github.com/gabrielsouzacoder/clean-new/entity"

type Todo struct {
	ID          entity.ID `json:"id"`
	Description string    `json:"description"`
	Status      bool      `json:"status"`
}

func (todo *Todo) ToTodo() *entity.Todo {
	return &entity.Todo{
		ID:          todo.ID,
		Status:      todo.Status,
		Description: todo.Description,
	}
}
