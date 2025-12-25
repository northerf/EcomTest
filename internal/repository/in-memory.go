package repository

import (
	"Ecom/internal/schema"
	"errors"
	"sync"
)

var (
	ErrNotFound = errors.New("todo not found")
)

type MemoryRepos struct {
	mu      sync.RWMutex
	todoMap map[int]schema.Todo
	nextID  int
}

func NewRepos() *MemoryRepos {
	return &MemoryRepos{
		todoMap: make(map[int]schema.Todo),
		nextID:  1,
	}
}

func (r *MemoryRepos) Create(todo schema.Todo) (schema.Todo, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	todo.ID = r.nextID
	r.todoMap[todo.ID] = todo
	r.nextID++
	return todo, nil
}

func (r *MemoryRepos) GetAll() []schema.Todo {
	r.mu.RLock()
	defer r.mu.RUnlock()
	todoMap := make([]schema.Todo, 0, len(r.todoMap))
	for _, todo := range r.todoMap {
		todoMap = append(todoMap, todo)
	}
	return todoMap
}

func (r *MemoryRepos) GetByID(id int) (schema.Todo, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	todo, ok := r.todoMap[id]
	if !ok {
		return schema.Todo{}, ErrNotFound
	}
	return todo, nil
}

func (r *MemoryRepos) Update(id int, todo schema.Todo) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.todoMap[id]; !ok {
		return ErrNotFound
	}
	todo.ID = id
	r.todoMap[id] = todo
	return nil
}

func (r *MemoryRepos) Delete(id int) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.todoMap[id]; !ok {
		return ErrNotFound
	}
	delete(r.todoMap, id)
	return nil
}
