package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"todo_golang/models"
)

var DB *gorm.DB

func Connect() {
	db, err := gorm.Open(sqlite.Open("dev.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Erro ao conectar com SQLite:", err)
	}

	db.AutoMigrate(&models.Task{})
	DB = db
}
