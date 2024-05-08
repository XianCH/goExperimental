package repository

import "github.com/x14n/goExperimental/gin-crud-api/v2/models"

type TodoRepository interface {
	Create(todo *models.Todo) (*models.Todo, error)
	GetAll() ([]*models.Todo, error)
}
