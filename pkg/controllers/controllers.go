package controllers

import (
	models "dependency/pkg/moldels"
	"dependency/pkg/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// TaskController represents a controller for the tasks
type TaskController struct {
	service services.TaskService
}

// NewTaskController creates a new instance of TaskCntroller
func NewTaskController(service services.TaskService) *TaskController {
	return &TaskController{service}
}

// SetupTaskRoutes configurates the routes related to Tasks
func SetupTaskRoutes(router *gin.Engine, service services.TaskService) {
	controller := NewTaskController(service)

	//criar as funções antes das rotas
}

// CreateTask manipulator to create a new task
func (c *TaskController) CreateTask(ctx *gin.Context) {
	var task models.Task
	if err := ctx.ShouldBindJSON(&task); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create task"})
		return
	}
	ctx.JSON(http.StatusCreated, task)
}

// GetTasks manipulator to Get all tasks
func (c *TaskController) GetTasks(ctx *gin.Context) {
	tasks, err := c.service.GetTasks()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed tp fetch all tasks"})
		return
	}
	ctx.JSON(http.StatusOK, tasks)
}

// GetTaskByID manipulator to get new task by ID
func (c *TaskController) GetTaskByID(ctx *gin.Context) {
	idStrg := ctx.Param("id")
	id, err := strconv.Atoi(idStrg)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}
	task, err := c.service.GetTaskByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch task"})
		return
	}
	if task == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	ctx.JSON(http.StatusOK, task)
}

// UpdateTask manipulator to update a task
func (c *TaskController) UpdateTask(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}
	var updatedTask models.Task
	if err := ctx.ShouldBindJSON(&updatedTask); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedTask.ID = uint(id)

	if err := c.service.UpdateTask(&updatedTask); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update task"})
		return
	}
	ctx.JSON(http.StatusOK, updatedTask)
}

// DeleteTask manipulator to delete task
func (c *TaskController) DeleteTask(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}
	if err := c.service.DeleteTask(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete task"})
		return
	}
	ctx.Status(http.StatusNoContent)
}
