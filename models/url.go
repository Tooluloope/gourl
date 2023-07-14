package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type URL struct {
	gorm.Model
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	OriginalURL string    `gorm:"type:varchar(2048)" json:"originalURL"`
	ShortURL    string    `gorm:"type:varchar(255);uniqueIndex" json:"shortURL"`
	UserID      uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"-"`
	User        User      `gorm:"foreignKey:UserID" json:"-"`
}
