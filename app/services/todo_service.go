package services

import (
	"github.com/johnnyFR26/go.api/app/models"
	"github.com/johnnyFR26/go.api/config"
)

func GetAllTodos() ([]models.Todo, error) {
	var todos []models.Todo
	err := config.DB.Find(&todos).Error
	return todos, err
}

func CreateTodo(todo *models.Todo) error {
	return config.DB.Create(todo).Error
}

func DeleteTodo(id string) error {
	return config.DB.Delete(&models.Todo{}, id).Error
}
