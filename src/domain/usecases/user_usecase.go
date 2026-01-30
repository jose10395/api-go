package usecases

import (
	"api-go/src/domain/entities"
	repositories "api-go/src/domain/ports"
	"errors"
)

type UserUsecase struct {
	userRepository repositories.UserRepository
}

func NewUserUsecase(userRepository repositories.UserRepository) *UserUsecase {
	return &UserUsecase{
		userRepository: userRepository,
	}
}

func (s *UserUsecase) Create(user *entities.User) error {
	if user.Name == "" {
		return errors.New("nombre requerido")
	}
	return s.userRepository.Create(user)
}

func (s *UserUsecase) FindByID(id uint) (*entities.User, error) {
	if id == 0 {
		return nil, errors.New("id inválido")
	}
	return s.userRepository.FindByID(id)
}

func (s *UserUsecase) FindAll() ([]entities.User, error) {
	return s.userRepository.FindAll()
}
