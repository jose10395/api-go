package repositories

import (
	"api-go/src/data/models"
	"api-go/src/domain/entities"

	"gorm.io/gorm"
)

type UserGormRepository struct {
	db *gorm.DB
}

func NewUserGormRepository(db *gorm.DB) *UserGormRepository {
	return &UserGormRepository{db: db}
}

func (r *UserGormRepository) Create(user *entities.User) error {
	model := models.UserModel{
		Name:  user.Name,
		Email: user.Email,
	}

	if err := r.db.Create(&model).Error; err != nil {
		return err
	}

	user.ID = model.ID
	return nil
}

func (r *UserGormRepository) FindAll() ([]entities.User, error) {
	var modelsList []models.UserModel

	if err := r.db.Find(&modelsList).Error; err != nil {
		return nil, err
	}

	users := make([]entities.User, 0, len(modelsList))
	for _, m := range modelsList {
		users = append(users, entities.User{
			ID:    m.ID,
			Name:  m.Name,
			Email: m.Email,
		})
	}

	return users, nil
}

func (r *UserGormRepository) FindByID(id uint) (*entities.User, error) {
	var model models.UserModel

	if err := r.db.First(&model, id).Error; err != nil {
		return nil, err
	}

	user := &entities.User{
		ID:    model.ID,
		Name:  model.Name,
		Email: model.Email,
	}

	return user, nil
}
