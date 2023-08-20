package models

import (
	"errors"

	"gorm.io/gorm"
)

type Pair struct {
	ID        uint `json:"id" gorm:"primaryKey"`
	Player1ID uint `json:"player1"`
	Player2ID uint `json:"player2"`
}

func (d *Pair) BeforeSave(tx *gorm.DB) (err error) {
	if d.Player1ID == d.Player2ID {
		return errors.New("the players cant be the same")
	}

	var count int64
	tx.Model(&Pair{}).
		Where("(Player1_id = ? AND Player2_id = ?) OR (Player1_id = ? AND Player2_id = ?)",
			d.Player1ID, d.Player2ID, d.Player2ID, d.Player1ID).
		Count(&count)

	if count > 0 {
		return errors.New("pair already exists")
	}

	return nil
}
