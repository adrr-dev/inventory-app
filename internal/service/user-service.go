package service

import (
	"fmt"

	"github.com/adrr-dev/inventory-app/internal/repository"
	"gorm.io/gorm"
)

type UserService struct {
	DB *gorm.DB
}

func (s UserService) FetchUser(username, password string) (*repository.User, error) {
	var user repository.User
	result := s.DB.Where("username = ? AND password = ?", username, password).First(&user)
	if result.Error != nil {
		return nil, fmt.Errorf("user not found, username or password incorrect: %e", result.Error)
	}

	return &user, nil
}

func (s UserService) CreateUser(username, password string) error {
	newUser := &repository.User{Username: username, Password: password}
	result := s.DB.Create(newUser)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s UserService) DeleteUser(username, password string) error {
	user, err := s.FetchUser(username, password)
	if err != nil {
		return err
	}

	result := s.DB.Where("id = ?", user.ID).Delete(&repository.User{})
	if result.Error != nil {
		return result.Error
	}

	// errors returned will most likely be errors about non existing user

	return nil
}
