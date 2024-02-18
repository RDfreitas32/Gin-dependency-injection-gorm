package services

import (
	models "dependency/pkg/moldels"
	"dependency/pkg/repositories"
)

// TaskServives implements  the logic of businnes to tasks
type TaskServive struct {
	repository repositories.TaskRepository
}

// NewTaskService creates a new instance of TaskService
func NewTaskService(repository repositories.TaskRepository) *TaskServive {
	return &TaskServive{repository}
}

// CreateTasks creats a new task
func (s *TaskServive) GetTasks() ([]models.Task, error) {
	return s.repository.GetAll()
}

// GetTaskByID returns a tasks by ID
func (s *TaskService) GetTaskByID(id uint) (*models.Task, error) {
	return s.repository.GetTaskByID{id}
}
