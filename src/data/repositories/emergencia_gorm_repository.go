package repositories

import (
	"api-go/src/data/models"
	"api-go/src/domain/entities"

	"gorm.io/gorm"
)

type EmergenciaGormRepository struct {
	db *gorm.DB
}

func NewEmergenciaGormRepository(db *gorm.DB) *EmergenciaGormRepository {
	return &EmergenciaGormRepository{db: db}
}

func (r *EmergenciaGormRepository) Create(e *entities.Emergencia) error {
	model := models.EmergenciaModel{
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

func (r *EmergenciaGormRepository) FindAll() ([]entities.Emergencia, error) {
	var modelsList []models.EmergenciaModel

	if err := r.db.Find(&modelsList).Error; err != nil {
		return nil, err
	}

	emergencias := make([]entities.Emergencia, 0, len(modelsList))
	for _, m := range modelsList {
		emergencias = append(emergencias, entities.Emergencia{
			ID_EMERGENCIA: m.ID_EMERGENCIA,
			Descripcion:   m.Descripcion,
			Fecha:         m.Fecha,
			Estado:        m.Estado,
		})
	}

	return emergencias, nil
}

func (r *EmergenciaGormRepository) FindByID(id uint) (*entities.Emergencia, error) {
	var model models.EmergenciaModel

	if err := r.db.First(&model, id).Error; err != nil {
		return nil, err
	}

	emergencia := &entities.Emergencia{
		ID_EMERGENCIA: model.ID_EMERGENCIA,
		Descripcion:   model.Descripcion,
		Fecha:         model.Fecha,
		Estado:        model.Estado,
	}

	return emergencia, nil
}
