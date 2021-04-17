package todo

import "github.com/gabrielsouzacoder/clean-new/entity"

type InMemory struct {
	m map[entity.ID]*entity.Todo
}

func NewInMemoryDatabase() *InMemory {
	var m = map[entity.ID]*entity.Todo{}

	return &InMemory{
		m: m,
	}
}

func (r *InMemory) Create(e *entity.Todo) (*entity.ID, error) {
	r.m[e.ID] = e
	return &e.ID, nil
}

func (r *InMemory) List() ([]*entity.Todo, error) {
	var d []*entity.Todo
	for _, j := range r.m {
		d = append(d, j)
	}

	return d, nil
}
