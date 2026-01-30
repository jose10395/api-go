package ports

import "api-go/src/domain/entities"

type UserRepository interface {
	Create(user *entities.User) error
	FindAll() ([]entities.User, error)
	FindByID(id uint) (*entities.User, error)
}
