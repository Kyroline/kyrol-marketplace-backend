package model

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	ID        uint       `gorm:"primaryKey;not null"`
	Name      string     `gorm:"type:varchar(255); not null"`
	Products  []*Product `gorm:"many2many:category_product"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (entity *Category) BeforeCreate(db *gorm.DB) error {
	entity.CreatedAt = time.Now().Local()
	return nil
}

func (entity *Category) BeforeUpdate(db *gorm.DB) error {
	entity.UpdatedAt = time.Now().Local()
	return nil
}
