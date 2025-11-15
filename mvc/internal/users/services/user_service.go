package services

import (
	"library-mvc/internal/users/models"
	"time"
)

type UserService struct {
	userRepository models.UserRepository
}

func NewUserService(userRepository models.UserRepository) models.UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (u *UserService) CreateUser(user *models.User) error {
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	return u.userRepository.CreateUser(user)
}

func (u *UserService) DeleteUser(id int64) error {
	return u.userRepository.DeleteUser(id)
}

func (u *UserService) GetAllUsers() ([]*models.User, error) {
	return u.userRepository.GetAllUsers()
}

func (u *UserService) GetUser(id int64) (*models.User, error) {
	return u.userRepository.GetUser(id)
}

func (u *UserService) UpdateUser(id int64, user *models.User) error {
	user.UpdatedAt = time.Now()
	return u.userRepository.UpdateUser(id, user)
}
