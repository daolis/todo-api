package data

import (
	"fmt"

	"github.com/google/uuid"
)

type Todo struct {
	ID          string `json:"_id"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

var Todos *TodosData

type TodosData struct {
	index map[string]int
	data  []*Todo
}

func (t *TodosData) GetTodoById(id string) (*Todo, error) {
	if index, exists := t.index[id]; exists {
		result := *t.data[index]
		return &result, nil
	}
	return nil, fmt.Errorf("todo with id '%s' not found", id)
}

func (t *TodosData) GetTodos() []*Todo {
	return t.data
}

func (t *TodosData) CreateTodo(description string) *Todo {
	idx := t.nextIndex()
	newTodo := &Todo{
		ID:          idx,
		Description: description,
		Done:        false,
	}
	t.data = append(t.data, newTodo)
	t.index[idx] = len(t.data) - 1
	return newTodo
}

func (t *TodosData) SetDone(id string) (*Todo, error) {
	if index, exists := t.index[id]; exists {
		t.data[index].Done = true
		result := *t.data[index]
		return &result, nil
	}
	return nil, fmt.Errorf("todo with id '%s' not found", id)
}

func (t *TodosData) nextIndex() string {
	return uuid.New().String()
}

func init() {
	Todos = &TodosData{
		data:  []*Todo{},
		index: make(map[string]int),
	}

	// initial data
	Todos.CreateTodo("test todo 01")
	Todos.CreateTodo("test todo 02")
}
