package models

import "time"

type EmergenciaModel struct {
	ID_EMERGENCIA uint
	Descripcion   string
	Fecha         time.Time
	Estado        string
}

func (EmergenciaModel) TableName() string {
	return "emergencia"
}
