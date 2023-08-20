package models

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type Player struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"unique"`
}

func (d *Player) IsPlayerExists(tx *gorm.DB) (err error) {
	//d.ID

	//var count int64
	//tx.Model(&Player{}).
	//	Where("(Player1_id = ? AND Player2_id = ?) OR (Player1_id = ? AND Player2_id = ?)",
	//		d.Player1ID, d.Player2ID, d.Player2ID, d.Player1ID).
	//	Count(&count)
	if err := tx.Model(&Player{}).First(&d, d.ID).Error; err != nil {
		e := fmt.Sprintf("Player %d not found", d.ID)
		return errors.New(e)
	}

	//if count > 0 {
	//	return errors.New("pair already exists")
	//}

	return nil
}
