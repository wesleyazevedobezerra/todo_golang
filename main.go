package main

import (
	"github.com/gin-gonic/gin"
	"todo_golang/controllers"
	"todo_golang/database"
)

func main() {
	database.Connect()
	r := gin.Default()

	api := r.Group("/api")
	{
		api.GET("/tasks", controllers.GetTasks)
		api.GET("/tasks/:id", controllers.GetTaskByID)
		api.POST("/tasks", controllers.CreateTask)
		api.PUT("/tasks/:id", controllers.UpdateTask)
		api.DELETE("/tasks/:id", controllers.DeleteTask)
	}

	r.Run(":8080") // Inicia em http://localhost:8080
}
