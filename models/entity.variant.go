package model

import (
	"time"

	"gorm.io/gorm"
)

type ProductVariant struct {
	ProductID string  `gorm:"type:varchar(15);not null"`
	ID        string  `gorm:"primaryKey"`
	Name      string  `gorm:"type:varchar(255);not null"`
	Price     float64 `gorm:"not null"`
	Stock     uint    `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (entity *ProductVariant) BeforeCreate(db *gorm.DB) error {
	entity.CreatedAt = time.Now().Local()
	return nil
}

func (entity *ProductVariant) BeforeUpdate(db *gorm.DB) error {
	entity.UpdatedAt = time.Now().Local()
	return nil
}
