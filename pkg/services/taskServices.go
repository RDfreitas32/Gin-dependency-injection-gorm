package services

import (
	models "dependency/pkg/moldels"
	"dependency/pkg/repositories"
)

// TaskServives implements  the logic of businnes to tasks
type TaskService struct {
	repository repositories.TaskRepository
}

// NewTaskService creates a new instance of TaskService
func NewTaskService(repository repositories.TaskRepository) *TaskService {
	return &TaskService{repository}
}

// CreateTasks creats a new task
func (s *TaskService) GetTasks() ([]models.Task, error) {
	return s.repository.GetAll()
}

// GetTaskByID returns a tasks by ID
func (s *TaskService) GetTaskByID(id uint) (*models.Task, error) {
	return s.repository.GetByID(id)
}

// UpdateTask updates a task
func (s *TaskService) UpdateTask(task *models.Task) error {
	return s.repository.Update(task)
}

// DeleteTask delets a task by ID
func (s *TaskService) DeleteTask(id uint) error {
	return s.repository.Delete(id)
}
