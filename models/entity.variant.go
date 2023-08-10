package model

import (
	"time"

	"gorm.io/gorm"
)

type EntityVariants struct {
	ProductID string  `gorm:"type:varchar(15);not null"`
	ID        string  `gorm:"primaryKey"`
	Name      string  `gorm:"type:varchar(255);not null"`
	Price     float64 `gorm:"not null"`
	Stock     uint    `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time `gorm:"default:null"`
}

func (entity *EntityVariants) BeforeCreate(db *gorm.DB) error {
	entity.CreatedAt = time.Now().Local()
	return nil
}

func (entity *EntityVariants) BeforeUpdate(db *gorm.DB) error {
	entity.UpdatedAt = time.Now().Local()
	return nil
}
