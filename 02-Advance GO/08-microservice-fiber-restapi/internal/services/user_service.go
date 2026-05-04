package services

import (
	"08-microservice-fiber-restapi/internal/models"
	"08-microservice-fiber-restapi/internal/repositories"
)

type UserService struct {
	userRepository *repositories.UserRepository
}

func NewUserService(r *repositories.UserRepository) *UserService {
	return &UserService{userRepository: r}
}

func (s *UserService) GetUserByID(id int) (*models.User, error) {
	return s.userRepository.FindByID(int64(id))
}

func (s *UserService) GetUserByName(name string) (*models.User, error) {
	return s.userRepository.FindByName(name)
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
	return s.userRepository.FindAllUsers()
}

func (s *UserService) CreateUser(user *models.User) error {
	return s.userRepository.CreateUser(user)
}

func (s *UserService) UpdateUser(user *models.User) error {
	return s.userRepository.UpdateUser(user)
}

func (s *UserService) DeleteUser(user *models.User) error {
	return s.userRepository.DeleteUser(user)
}
