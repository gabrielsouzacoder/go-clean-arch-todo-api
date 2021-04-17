package entity_test

import (
	"errors"
	"github.com/gabrielsouzacoder/clean-new/entity"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewTodo(t *testing.T) {
	b, err := entity.NewTodo("Read a book", false)
	assert.Nil(t, err)
	assert.Equal(t, b.Description, "Read a book")
	assert.Equal(t, b.Status, false)
}

func TestInvalidTodoDescription(t *testing.T) {
	b, err := entity.NewTodo("", false)

	assert.Nil(t, b)
	assert.Equal(t, err, errors.New("invalid entity"))
}