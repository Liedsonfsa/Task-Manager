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
	query := `INSERT INTO tasks (title, description, status) VALUES (?, ?, ?)`
	result, err := r.db.Exec(query, task.Title, task.Description, task.Status)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	task.ID = int(id)

	row := r.db.QueryRow("SELECT created_at, updated_at FROM tasks WHERE id = ?", task.ID)
	err = row.Scan(&task.CreatedAt, &task.UpdatedAt)

	return err
}

func (r *TaskRepository) GetTaskByID(id int) (*models.Task, error) {
	task := &models.Task{}
	query := `SELECT id, title, description, status, created_at, updated_at FROM tasks WHERE id = ?`

	err := r.db.QueryRow(query, id).Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.CreatedAt, &task.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (r *TaskRepository) UpdateTask(task *models.Task) error {
	query := `UPDATE tasks SET title=?, description=?, status=? WHERE id=?`
	_, err := r.db.Exec(query, task.Title, task.Description, task.Status, task.ID)
	if err != nil {
		return err
	}

	return r.db.QueryRow("SELECT updated_at FROM tasks WHERE id = ?", task.ID).Scan(&task.UpdatedAt)
}

func (r *TaskRepository) DeleteTask(id int) error {
	query := `DELETE FROM tasks WHERE id=?`
	_, err := r.db.Exec(query, id)
	return err
}

func (r *TaskRepository) GetAllTasks() ([]models.Task, error) {
	query := `SELECT id, title, description, status, created_at, updated_at FROM tasks`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var task models.Task

		if err = rows.Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.CreatedAt, &task.UpdatedAt); err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}