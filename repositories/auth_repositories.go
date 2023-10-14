package repositories

import (
	"learn-echo/configs"
	"learn-echo/models"
)

func Register(user *models.User) error {
	result := configs.DB.Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetUserByEmail(email string) (*models.User, error) {
	user := &models.User{}
	result := configs.DB.Where("email = ?", email).First(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}
