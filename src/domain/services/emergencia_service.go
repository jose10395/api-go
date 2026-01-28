package services

import (
	"api-go/src/domain/entities"
	"api-go/src/domain/repositories"
	"errors"
)

type EmergenciaService struct {
	emergenciaRepository repositories.EmergenciaRepository
}

func NewEmergenciaService(emergenciaRepository repositories.EmergenciaRepository) *EmergenciaService {
	return &EmergenciaService{
		emergenciaRepository: emergenciaRepository,
	}
}

func (s *EmergenciaService) Create(emergencia *entities.Emergencia) error {
	if emergencia.Descripcion == "" {
		return errors.New("Descripcion requerida")
	}
	return s.emergenciaRepository.Create(emergencia)
}

func (s *EmergenciaService) FindByID(id uint) (*entities.Emergencia, error) {
	if id == 0 {
		return nil, errors.New("id inválido")
	}
	return s.emergenciaRepository.FindByID(id)
}

func (s *EmergenciaService) FindAll() ([]entities.Emergencia, error) {
	return s.emergenciaRepository.FindAll()
}
