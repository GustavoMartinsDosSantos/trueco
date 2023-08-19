package models

import "time"

type Partida struct {
	ID         uint `gorm:"primaryKey"`
	Dupla1ID   uint
	Dupla2ID   uint
	VencedorID uint
	Data       time.Time
}
