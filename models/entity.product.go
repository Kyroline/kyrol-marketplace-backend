package model

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID             string           `gorm:"primaryKey;type:varchar(15)"`
	Name           string           `gorm:"type:varchar(255);not null"`
	Description    string           `gorm:"type:varchar(255);not null"`
	ProductVariant []ProductVariant `gorm:"foreignKey:ProductID;references:ID"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
}

func (entity *Product) BeforeCreate(db *gorm.DB) error {
	entity.CreatedAt = time.Now().Local()
	return nil
}

func (entity *Product) BeforeUpdate(db *gorm.DB) error {
	entity.UpdatedAt = time.Now().Local()
	return nil
}
