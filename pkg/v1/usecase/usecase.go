package usecase

import (
	"errors"

	"github.com/zerokkcoder/grpc-clean/internal/models"
	interfaces "github.com/zerokkcoder/grpc-clean/pkg/v1"
	"gorm.io/gorm"
)

type UseCase struct {
	repo interfaces.RepoInterface
}

func New(repo interfaces.RepoInterface) interfaces.UseCaseInterface {
	return &UseCase{repo}
}

// Create
//
// This function creates a user with the data supplied
func (uc *UseCase) Create(user models.User) (models.User, error) {
	// check if valid email is supplied
	if _, err := uc.repo.GetByEmail(user.Email); !errors.Is(err, gorm.ErrRecordNotFound) {
		return models.User{}, errors.New("the email is already associated with another user")
	}
	// email doesnot exist so,now proceed
	return uc.repo.Create(user)
}

// Get
//
// This function returns the user instance
func (uc *UseCase) Get(id string) (models.User, error) {
	return uc.repo.Get(id)
}

// Update
//
// This function updates the user and returns if any error occurred
func (uc *UseCase) Update(user models.User) error {
	return uc.repo.Update(user)
}

// Delete
//
// This function deletes the user and returns if any error occurred
func (uc *UseCase) Delete(id string) error {
	return uc.repo.Delete(id)
}
