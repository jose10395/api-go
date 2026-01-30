package services

import (
	"api-go/src/domain/entities"
	"api-go/src/domain/repositories"
	"errors"
)

type EmergencyService struct {
	EmergencyRepository repositories.EmergencyRepository
}

func NewEmergencyService(EmergencyRepository repositories.EmergencyRepository) *EmergencyService {
	return &EmergencyService{
		EmergencyRepository: EmergencyRepository,
	}
}

func (s *EmergencyService) Create(emergency *entities.Emergency) error {
	if emergency.Descripcion == "" {
		return errors.New("Descripcion requerida")
	}
	return s.EmergencyRepository.Create(emergency)
}

func (s *EmergencyService) FindByID(id uint) (*entities.Emergency, error) {
	if id == 0 {
		return nil, errors.New("id inválido")
	}
	return s.EmergencyRepository.FindByID(id)
}

func (s *EmergencyService) FindAll() ([]entities.Emergency, error) {
	return s.EmergencyRepository.FindAll()
}
