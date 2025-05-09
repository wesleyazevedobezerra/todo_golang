package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "todo_golang/database"
    "todo_golang/models"
)

func GetTasks(c *gin.Context) {
    var tasks []models.Task
    database.DB.Find(&tasks)
    c.JSON(http.StatusOK, tasks)
}

func GetTaskByID(c *gin.Context) {
    id := c.Param("id")
    var task models.Task
    if err := database.DB.First(&task, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Tarefa não encontrada"})
        return
    }
    c.JSON(http.StatusOK, task)
}

func CreateTask(c *gin.Context) {
    var task models.Task
    if err := c.ShouldBindJSON(&task); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    database.DB.Create(&task)
    c.JSON(http.StatusCreated, task)
}

func UpdateTask(c *gin.Context) {
    id := c.Param("id")
    var task models.Task
    if err := database.DB.First(&task, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Tarefa não encontrada"})
        return
    }
    var input models.Task
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    task.Titulo = input.Titulo
    task.Descricao = input.Descricao
    task.Status = input.Status
    database.DB.Save(&task)
    c.JSON(http.StatusOK, task)
}

func DeleteTask(c *gin.Context) {
    id := c.Param("id")
    database.DB.Delete(&models.Task{}, id)
    c.Status(http.StatusNoContent)
}