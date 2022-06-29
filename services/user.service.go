package services

import "github.com/TomasHansut/gin_api/models"

// Interface that will help to define api contracts
type UserService interface {
	// Pass User obejct and return error
	CreateUser(*models.User) error
	// Pass user name as string and return User object or error
	GetUser(*string) (*models.User, error)
	// Return slice of users or error
	GetAll() ([]*models.User, error)
	// Pass user object return error
	UpdateUser(*models.User) error
	// Pass user name as string return error
	DeleteUser(*string) error
}
