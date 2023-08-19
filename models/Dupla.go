package models

import (
	"errors"

	"gorm.io/gorm"
)

type Dupla struct {
	ID         uint `gorm:"primaryKey"`
	Jogador1ID uint
	Jogador2ID uint
}

func (d *Dupla) BeforeSave(tx *gorm.DB) (err error) {
	if d.Jogador1ID == d.Jogador2ID {
		return errors.New("jogadores não podem ser iguais em uma dupla")
	}

	var count int64
	tx.Model(&Dupla{}).
		Where("(jogador1_id = ? AND jogador2_id = ?) OR (jogador1_id = ? AND jogador2_id = ?)",
			d.Jogador1ID, d.Jogador2ID, d.Jogador2ID, d.Jogador1ID).
		Count(&count)

	if count > 0 {
		return errors.New("já existe uma dupla com esses jogadores")
	}

	return nil
}
