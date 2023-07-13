package models

import "gorm.io/gorm"

type URL struct {
	gorm.Model
	OriginalURL string `gorm:"type:varchar(2048)"`
	ShortURL    string `gorm:"type:varchar(255);uniqueIndex"`
	UserID      uint
}
