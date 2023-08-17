package model

import (
	"time"

	"gorm.io/gorm"
)

type Cart struct {
	UserID    string
	User      User       `gorm:"foreignKey:UserID"`
	ID        string     `gorm:"primaryKey"`
	Item      []CartItem `gorm:"foreignKey:CartID"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (entity *Cart) BeforeCreate(db *gorm.DB) error {
	entity.CreatedAt = time.Now().Local()
	return nil
}

func (entity *Cart) BeforeUpdate(db *gorm.DB) error {
	entity.UpdatedAt = time.Now().Local()
	return nil
}
