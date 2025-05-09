package main

import (
	"github.com/gin-gonic/gin"
	"todo_golang/controllers"
	"todo_golang/database"
)

func main() {
	// Conectar ao banco de dados
	database.Connect()
	r := gin.Default()

	// Definindo as rotas
	tasks := "/tasks"       // Rota para todas as tarefas
	tasksId := "/tasks/:id" // Rota para tarefa específica com ID

	api := r.Group("/api")
	{
		// Rota para obter todas as tarefas
		api.GET(tasks, controllers.GetTasks)

		// Rota para obter uma tarefa específica
		api.GET(tasksId, controllers.GetTaskByID)

		// Rota para criar uma nova tarefa
		api.POST(tasks, controllers.CreateTask)

		// Rota para atualizar uma tarefa específica
		api.PUT(tasksId, controllers.UpdateTask)

		// Rota para deletar uma tarefa específica
		api.DELETE(tasksId, controllers.DeleteTask)
	}

	// Iniciar o servidor
	r.Run(":8080") // Inicia em http://localhost:8080
}
