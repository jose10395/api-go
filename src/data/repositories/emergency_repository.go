package repositories

import (
	"api-go/src/data/models"
	"api-go/src/domain/entities"

	"gorm.io/gorm"
)

type EmergencyRepositoryImpl struct {
	db *gorm.DB
}

func NewEmergencyRepositoryImpl(db *gorm.DB) *EmergencyRepositoryImpl {
	return &EmergencyRepositoryImpl{db: db}
}

func (r *EmergencyRepositoryImpl) Create(e *entities.Emergency) error {
	model := models.EmergencyModel{
		Descripcion: e.Descripcion,
		Fecha:       e.Fecha,
		Estado:      e.Estado,
	}

	if err := r.db.Create(&model).Error; err != nil {
		return err
	}

	e.ID_EMERGENCIA = model.ID_EMERGENCIA
	return nil
}

func (r *EmergencyRepositoryImpl) FindAll() ([]entities.Emergency, error) {
	var modelsList []models.EmergencyModel

	if err := r.db.Find(&modelsList).Error; err != nil {
		return nil, err
	}

	emergencies := make([]entities.Emergency, 0, len(modelsList))
	for _, m := range modelsList {
		emergencies = append(emergencies, entities.Emergency{
			ID_EMERGENCIA: m.ID_EMERGENCIA,
			Descripcion:   m.Descripcion,
			Fecha:         m.Fecha,
			Estado:        m.Estado,
		})
	}

	return emergencies, nil
}

func (r *EmergencyRepositoryImpl) FindByID(id uint) (*entities.Emergency, error) {
	var model models.EmergencyModel

	if err := r.db.First(&model, id).Error; err != nil {
		return nil, err
	}

	emergency := &entities.Emergency{
		ID_EMERGENCIA: model.ID_EMERGENCIA,
		Descripcion:   model.Descripcion,
		Fecha:         model.Fecha,
		Estado:        model.Estado,
	}

	return emergency, nil
}
