package entities

import (
	"time"
)

type Emergencia struct {
	ID_EMERGENCIA uint
	Descripcion   string
	Fecha         time.Time
	Estado        string
}
