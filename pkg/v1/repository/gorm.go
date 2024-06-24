package repository

import (
	"github.com/zerokkcoder/grpc-clean/internal/models"
	interfaces "github.com/zerokkcoder/grpc-clean/pkg/v1"
	"gorm.io/gorm"
)

type Repo struct {
	db *gorm.DB
}

func New(db *gorm.DB) interfaces.RepoInterface {
	return &Repo{db}
}

// Create
//
// This function creates a new User which was supplied as the argument
func (repo *Repo) Create(user models.User) (models.User, error) {
	err := repo.db.Create(&user).Error
	return user, err
}

// Get
//
// This function returns the user instance which is
// saved on the DB and returns to the usecase
func (repo *Repo) Get(id string) (models.User, error) {
	var user models.User
	err := repo.db.Where("id = ?", id).First(&user).Error

	return user, err
}

// Update
//
// This function updates the user and returns if any error occurred
func (repo *Repo) Update(user models.User) error {
	return repo.db.Updates(&user).Error
}

// Delete
//
// This function deletes the user and returns if any error occurred
func (repo *Repo) Delete(id string) error {
	return repo.db.Where("id = ?", id).Delete(&models.User{}).Error
}

// GetByEmail
//
// This function returns the user instance which is
// saved on the DB and returns to the usecase
func (repo *Repo) GetByEmail(email string) (models.User, error) {
	var user models.User
	err := repo.db.Where("email = ?", email).First(&user).Error

	return user, err
}
