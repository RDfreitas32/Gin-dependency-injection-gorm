package main

import (
	"dependency/pkg/controllers"
	"dependency/pkg/db"
	"dependency/pkg/repositories"
	"dependency/pkg/services"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := db.ConectDB()
	if err != nil {
		panic(fmt.Errorf("Error to Conect DB: %w", err))
	}
	// Starting touter
	router := gin.Default()

	//Routes configuration
	taskRepository := repositories.NewTaskRepository(db)
	taskService := services.NewTaskService(taskRepository)
	controllers.SetupTaskRoutes(router, *taskService)

	//Executing service
	router.Run(":5000")
}
