package controllers

import (
	"dependency/pkg/services"

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
