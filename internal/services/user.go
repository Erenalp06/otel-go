package services

import (
	"github.com/Erenalp06/otel-go/internal/repository"
	"github.com/Erenalp06/otel-go/pkg/models"
)

type UserService struct {
	UserRepository *repository.UserRepository
}

type UserServiceInterface interface {
	CreateUser(user models.User) (models.User, error)
	GetUserById(id uint) (models.User, error)
	GetAllUsers() ([]models.User, error)
	UpdateUser(user models.User) (models.User, error)
	DeleteUser(id uint) error
}

func NewUserService(repository *repository.UserRepository) *UserService {
	return &UserService{
		UserRepository: repository,
	}
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
	return s.UserRepository.GetAllUsers()
}

func (s *UserService) CreateUser(user models.User) (models.User, error) {
	return s.UserRepository.CreateUser(user)
}

func (s *UserService) GetUserById(id uint) (models.User, error) {
	return s.UserRepository.GetUserById(id)
}

func (s *UserService) UpdateUser(user models.User) (models.User, error) {
	return s.UserRepository.UpdateUser(user)
}

func (s *UserService) DeleteUser(id uint) error {
	return s.UserRepository.DeleteUser(id)
}
