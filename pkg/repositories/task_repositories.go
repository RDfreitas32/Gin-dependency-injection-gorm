package repositories

import (
	models "dependency-gin/pkg/moldels"

	"gorm.io/gorm"
)

// TaskRepository defines methods to interact with tasks
type TaskRepository interface {
	Create(task *models.Task) error
	GetAll() ([]models.Task, error)
	GetByID(id uint) (*models.Task, error)
	Update(task *models.Task) error
	Delete(id uint) error
}

// taskRepository implements TaskRepository
type taskRepository struct {
	db *gorm.DB
}

// NewTaskRepository cries a new instance from taskRepository
func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{db}
}

// Create a new task
func (r *taskRepository) Create(task *models.Task) error {
	return r.db.Create(task).Error
}

// GetAll return all tasks
func (r *taskRepository) GetAll() ([]models.Task, error) {
	var tasks []models.Task
	if err := r.db.Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

// GetByID returns a task by ID
func (r *taskRepository) GetByID(id uint) (*models.Task, error) {
	var task models.Task
	if err := r.db.First(&task, id).Error; err != nil {
		return nil, err
	}
	return &task, nil
}

// Update updates tasks
func (r *taskRepository) Update(task *models.Task) error {
	return r.db.Save(task).Error
}

// Delete deletes tasks
func (r *taskRepository) Delete(id uint) error {
	return r.db.Delete(&models.Task{}, id).Error
}
