package service

import (
	"github.com/x14n/goExperimental/gin-crud-api/v2/models"
	"github.com/x14n/goExperimental/gin-crud-api/v2/repository"
)

type TodoServiceImpl struct {
	TodoRepository repository.TodoRepository
}

func NewTodoServiceImple(todoRepository repository.TodoRepository) TodoService {
	return &TodoServiceImpl{
		TodoRepository: todoRepository,
	}
}

func (t TodoServiceImpl) CreateTodo(todo *models.Todo) (*models.Todo, error) {
	todo, error := t.TodoRepository.Create(todo)
	return todo, error
}

func (t TodoServiceImpl) GetAlltodos() ([]*models.Todo, error) {
	return nil, nil
}
