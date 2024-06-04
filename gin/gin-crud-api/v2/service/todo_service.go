package service

import "github.com/x14n/goExperimental/gin-crud-api/v2/models"

type TodoService interface {
	CreateTodo(todo *models.Todo) (*models.Todo, error)
	GetAlltodos() ([]*models.Todo, error)
}
