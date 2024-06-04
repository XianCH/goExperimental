package repository

import (
	"database/sql"

	"github.com/x14n/goExperimental/gin-crud-api/v2/models"
)

type TodoRepositoryImpl struct {
	db *sql.DB
}

func NewTodoRepositoryImpl(db *sql.DB) TodoRepository {
	return &TodoRepositoryImpl{db: db}
}

func (r *TodoRepositoryImpl) Create(todo *models.Todo) (*models.Todo, error) {
	query := "INSERT INTO todos (title,status) VALUES (?,?)"
	result, err := r.db.Exec(query, todo.Title, todo.Status)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	todo.ID = int(id)
	return todo, nil
}

func (r *TodoRepositoryImpl) GetAll() ([]*models.Todo, error) {
	query := "SELECT id,title,status FROM todos"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []*models.Todo
	for rows.Next() {
		var todo models.Todo
		if err := rows.Scan(&todo.ID, &todo.Title, &todo.Status); err != nil {
			return nil, err
		}
		todos = append(todos, &todo)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return todos, nil
}
