package repository

import (
	"github.com/Erenalp06/otel-go/pkg/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

type UserRepositoryInterface interface {
	CreateUser(user models.User) (models.User, error)
	GetUserById(id uint) (models.User, error)
	GetAllUsers() ([]models.User, error)
	UpdateUser(user models.User) (models.User, error)
	DeleteUser(id uint) error
}

func NewRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (r *UserRepository) GetAllUsers() ([]models.User, error) {
	var users []models.User
	result := r.DB.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func (r *UserRepository) CreateUser(user models.User) (models.User, error) {
	result := r.DB.Create(&user)
	if result.Error != nil {
		return models.User{}, result.Error
	}
	return user, nil
}

func (r *UserRepository) GetUserById(id uint) (models.User, error) {
	var user models.User
	result := r.DB.First(&user, id)
	if result.Error != nil {
		return models.User{}, result.Error
	}
	return user, nil
}

func (r *UserRepository) UpdateUser(user models.User) (models.User, error) {
	result := r.DB.Save(&user)
	if result.Error != nil {
		return models.User{}, result.Error
	}
	return user, nil
}

func (r *UserRepository) DeleteUser(id uint) error {
	result := r.DB.Delete(&models.User{}, id)
	return result.Error
}
