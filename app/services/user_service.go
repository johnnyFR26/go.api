package services

import (
	"github.com/johnnyFR26/go.api/app/models"
	"github.com/johnnyFR26/go.api/config"
)

func GetAllUsers() ([]models.User, error) {
	var users []models.User
	err := config.DB.Preload("Tasks").Find(&users).Error
	return users, err
}

func CreateUser(user *models.User) (models.User, error) {
	err := config.DB.Create(user).Error
	return *user, err
}
