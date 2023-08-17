package model

import (
	"time"

	"gorm.io/gorm"
)

type CartItem struct {
	CartID         string
	ID             string
	ProductID      string
	Product        Product `gorm:"foreignKey:ProductID"`
	VariantID      string
	ProductVariant ProductVariant `gorm:"foreignKey:VariantID"`
	Qty            uint
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func (entity *CartItem) BeforeCreate(db *gorm.DB) error {
	entity.CreatedAt = time.Now().Local()
	return nil
}

func (entity *CartItem) BeforeUpdate(db *gorm.DB) error {
	entity.UpdatedAt = time.Now().Local()
	return nil
}
