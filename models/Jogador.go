package models

type Jogador struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Nome string `json:"name" gorm:"unique"`
}
