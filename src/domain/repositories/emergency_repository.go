package repositories

import "api-go/src/domain/entities"

type EmergencyRepository interface {
	Create(emergencia *entities.Emergency) error
	FindAll() ([]entities.Emergency, error)
	FindByID(id uint) (*entities.Emergency, error)
}
