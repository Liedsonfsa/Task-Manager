package service

import (
	"errors"

	"github.com/Liedsonfsa/Task-Manager/internal/models"
	"github.com/Liedsonfsa/Task-Manager/internal/repository"
)

type TaskService struct {
	repo *repository.TaskRepository
}

func NewTaskService(repo *repository.TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) CreateTask(task *models.Task) error {
	if task.Title == "" {
		return errors.New("title is required")
	}

	return s.repo.CreateTask(task)
}

func (s *TaskService) GetTask(id int) (*models.Task, error) {
	return s.repo.GetTaskByID(id)
}

func (s *TaskService) UpdateTask(task *models.Task) error {
	existingTask, err := s.repo.GetTaskByID(task.ID)
	if err != nil {
		return err
	}

	if existingTask == nil {
		return errors.New("task not found")
	}

	return s.repo.UpdateTask(task)
}

func (s *TaskService) DeleteTask(id int) error {
	return s.repo.DeleteTask(id)
}

func (s *TaskService) GetAllTasks() ([]models.Task, error) {
	return s.repo.GetAllTasks()
}