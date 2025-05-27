package config

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/johnnyFR26/go.api/app/models"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(sqlite.Open("todos.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database!", err)
	}

	database.AutoMigrate(&models.Todo{}, &models.User{}, &models.Task{})

	DB = database
}
