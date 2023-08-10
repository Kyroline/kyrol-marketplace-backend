package model

import (
	"time"

	"gorm.io/gorm"
)

type EntityProducts struct {
	ID             string           `gorm:"primaryKey;type:varchar(15)"`
	Name           string           `gorm:"type:varchar(255);not null"`
	Description    string           `gorm:"type:varchar(255);not null"`
	EntityVariants []EntityVariants `gorm:"foreignKey:ProductID;references:ID"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      time.Time `gorm:"default:null"`
}

func (entity *EntityProducts) BeforeCreate(db *gorm.DB) error {
	entity.CreatedAt = time.Now().Local()
	return nil
}

func (entity *EntityProducts) BeforeUpdate(db *gorm.DB) error {
	entity.UpdatedAt = time.Now().Local()
	return nil
}
