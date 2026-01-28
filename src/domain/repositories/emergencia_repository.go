package repositories

import "api-go/src/domain/entities"

type EmergenciaRepository interface {
	Create(emergencia *entities.Emergencia) error
	FindAll() ([]entities.Emergencia, error)
	FindByID(id uint) (*entities.Emergencia, error)
}
