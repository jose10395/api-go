package entities

import (
	"time"
)

type Emergency struct {
	ID_EMERGENCIA uint
	Descripcion   string
	Fecha         time.Time
	Estado        string
}
