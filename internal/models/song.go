package models

import (
	"gorm.io/gorm"
)

type Song struct {
	gorm.Model
	FileName string `gorm:"not null"`
	Title    string
	Artist   string
	Album    string
}
