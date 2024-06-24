package v1

import "github.com/zerokkcoder/grpc-clean/internal/models"

// RepoInterface this is an interface for repo methods
type RepoInterface interface {
	// Create creates a user with the data supplied
	Create(models.User) (models.User, error)
	// Get get retrieves the user instance
	Get(id string) (models.User, error)
	// Update update method updates the user and returns if any error occurred
	Update(models.User) error
	// Delete the user whose ID id supplied
	Delete(id string) error
	// GetByEmail returns the user instance which is
	GetByEmail(email string) (models.User, error)
}

// UseCaseInterface this is an interface for usecase methods
type UseCaseInterface interface {
	// Create a user with the data supplied
	Create(models.User) (models.User, error)
	// Get get retrieves the user instance
	Get(id string) (models.User, error)
	// Update update method updates the user and returns if any error occurred
	Update(models.User) error
	// Delete the user whose ID id supplied
	Delete(id string) error
}
