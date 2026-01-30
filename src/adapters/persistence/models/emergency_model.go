package models

import "time"

type EmergencyModel struct {
	ID_EMERGENCIA uint
	Descripcion   string
	Fecha         time.Time
	Estado        string
}

func (EmergencyModel) TableName() string {
	return "emergencia"
}
