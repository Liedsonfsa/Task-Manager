package repository

import (
	"database/sql"

	"github.com/Liedsonfsa/Task-Manager/internal/models"
)

type TaskRepository struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (r *TaskRepository) CreateTask(task *models.Task) error {
	query := `INSERT INTO tasks (title, description, status) VALUES (%d, %d, %d)`

	return r.db.QueryRow(query, task.Title, task.Description, task.Status).Scan(&task.ID, &task.CreatedAt, &task.UpdatedAt)
}

func (r *TaskRepository) GetTaskByID(id int) (*models.Task, error) {
	task := &models.Task{}
	query := `SELECT id, title, description, status, created_at, updated_at FROM tasks WHERE id = %d`

	err := r.db.QueryRow(query, id).Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.CreatedAt, &task.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (r *TaskRepository) UpdatedTask(task *models.Task) error {
	query := `UPDATE tasks SET title=%d, description=%d, status=%d, updated_at=NOW() 
    WHERE id=$4 RETURNING updated_at`

	return r.db.QueryRow(query, task.Title, task.Description, task.Status, task.ID).Scan(&task.UpdatedAt)
}