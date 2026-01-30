package usecases

import (
	"api-go/src/domain/entities"
	repositories "api-go/src/domain/ports"
	"errors"
)

type EmergencyUsecase struct {
	EmergencyRepository repositories.EmergencyRepository
}

func NewEmergencyUsecase(EmergencyRepository repositories.EmergencyRepository) *EmergencyUsecase {
	return &EmergencyUsecase{
		EmergencyRepository: EmergencyRepository,
	}
}

func (s *EmergencyUsecase) Create(emergency *entities.Emergency) error {
	if emergency.Descripcion == "" {
		return errors.New("Descripcion requerida")
	}
	return s.EmergencyRepository.Create(emergency)
}

func (s *EmergencyUsecase) FindByID(id uint) (*entities.Emergency, error) {
	if id == 0 {
		return nil, errors.New("ID inválido")
	}
	return s.EmergencyRepository.FindByID(id)
}

func (s *EmergencyUsecase) FindAll() ([]entities.Emergency, error) {
	return s.EmergencyRepository.FindAll()
}
