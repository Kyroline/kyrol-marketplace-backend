package model

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID          string      `gorm:"primaryKey;type:varchar(15);"`
	Name        string      `gorm:"type:varchar(255);not null"`
	Description string      `gorm:"type:varchar(255);not null"`
	Categories  []*Category `gorm:"many2many:category_product"`
	Price       float64     `gorm:"not null"`
	Stock       uint        `gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

func (entity *Product) BeforeCreate(db *gorm.DB) error {
	entity.CreatedAt = time.Now().Local()
	return nil
}

func (entity *Product) BeforeUpdate(db *gorm.DB) error {
	entity.UpdatedAt = time.Now().Local()
	return nil
}
