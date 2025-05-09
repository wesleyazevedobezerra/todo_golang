package main

import (
	"github.com/gin-gonic/gin"
	"todo_golang/controllers"
	"todo_golang/database"
)

func main() {
	database.Connect()
	r := gin.Default()
	taskId := "/tasks/:id"

	api := r.Group("/api")
	{
		api.GET("/tasks", controllers.GetTasks)
		api.GET(taskId, controllers.GetTaskByID)
		api.POST("/tasks", controllers.CreateTask)
		api.PUT(taskId, controllers.UpdateTask)
		api.DELETE(taskId, controllers.DeleteTask)
	}

	r.Run(":8080") // Inicia em http://localhost:8080
}
