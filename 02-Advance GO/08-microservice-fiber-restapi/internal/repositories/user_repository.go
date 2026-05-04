package repositories

import (
	"08-microservice-fiber-restapi/internal/models"
	"08-microservice-fiber-restapi/pkg/database"
)

type UserRepository struct{}

func (r *UserRepository) FindByID(id int64) (*models.User, error) {
	var user models.User
	err := database.DB.First(&user, id).Error
	return &user, err
}

func (r *UserRepository) FindByName(name string) (*models.User, error) {
	var user models.User
	err := database.DB.Where("name = ?", name).First(&user).Error
	return &user, err
}

func (r *UserRepository) FindAllUsers() ([]models.User, error) {
	var users []models.User
	err := database.DB.Find(&users).Error
	return users, err
}

func (r *UserRepository) CreateUser(user *models.User) error {
	return database.DB.Create(user).Error
}

func (r *UserRepository) UpdateUser(user *models.User) error {
	return database.DB.Save(user).Error
}

func (r *UserRepository) DeleteUser(user *models.User) error {
	return database.DB.Delete(user).Error
}
