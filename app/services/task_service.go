package services

import (
	"github.com/johnnyFR26/go.api/app/models"
	"github.com/johnnyFR26/go.api/config"
)

func GetAllTasks() ([]models.Task, error) {
	var tasks []models.Task
	err := config.DB.Find(&tasks).Error
	return tasks, err
}

func CreateTask(task *models.Task) (models.Task, error) {
	err := config.DB.Create(task).Error
	return *task, err
}

func DeleteTask(id string) error {
	return config.DB.Delete(&models.Task{}, id).Error
}
