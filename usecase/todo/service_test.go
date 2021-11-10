package todo

import (
	"github.com/gabrielsouzacoder/clean-new/entity"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateTodo(t *testing.T) {
	repo := NewInMemoryDatabase()

	service := NewService(repo)

	a, err := service.CreateTodo("Read books")

	assert.Nil(t, err)
	assert.NotNil(t, a)
}

func TestCreateTodoWithError(t *testing.T) {
	repo := NewInMemoryDatabase()

	service := NewService(repo)

	a, err := service.CreateTodo("")

	assert.NotNil(t, err)
	assert.Nil(t, a)
}

func TestGetAllEmptyList(t *testing.T) {
	repo := NewInMemoryDatabase()

	service := NewService(repo)

	todos, _ := service.ListTodos()

	assert.Equal(t, 0, len(todos))
}

func TestGetAllTodos(t *testing.T) {
	repo := NewInMemoryDatabase()

	service := NewService(repo)

	_, err := service.CreateTodo("Read a book")

	if err != nil {
		return
	}

	_, err = service.CreateTodo("Write a book")

	if err != nil {
		return
	}

	todos, err := service.ListTodos()

	assert.Equal(t, 2, len(todos))
}

func TestDeleteTodo(t *testing.T) {
	repo := NewInMemoryDatabase()

	service := NewService(repo)

	data, _ := service.CreateTodo("Read a book")

	service.DeleteTodo(data)

	todos, _ := service.ListTodos()

	assert.Equal(t, 0, len(todos))
}

func TestFindById(t *testing.T) {
	repo := NewInMemoryDatabase()

	service := NewService(repo)

	data, _ := service.CreateTodo("Read a book")

	todoById := service.FindById(data)

	assert.Equal(t, todoById.Description, "Read a book")
}

func TestFindByIdNotFound(t *testing.T) {
	repo := NewInMemoryDatabase()

	service := NewService(repo)

	id := entity.NewID()

	todoById := service.FindById(&id)

	assert.Nil(t, todoById)
}

func TestUpdateTodo(t *testing.T) {
	repo := NewInMemoryDatabase()

	service := NewService(repo)

	idCreated, _ := service.CreateTodo("Test of Todo")

	created := service.FindById(idCreated)

	created.Description = "New Description"

	todoUpdated := service.UpdateTodo(created)

	assert.Equal(t, todoUpdated.Description, "New Description")
}
