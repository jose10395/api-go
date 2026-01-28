package services

import (
	"api-go/src/domain/entities"
	"api-go/src/domain/repositories"
	"errors"
)

type UserService struct {
	userRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (s *UserService) Create(user *entities.User) error {
	if user.Name == "" {
		return errors.New("nombre requerido")
	}
	return s.userRepository.Create(user)
}

func (s *UserService) FindByID(id uint) (*entities.User, error) {
	if id == 0 {
		return nil, errors.New("id inválido")
	}
	return s.userRepository.FindByID(id)
}

func (s *UserService) FindAll() ([]entities.User, error) {
	return s.userRepository.FindAll()
}
