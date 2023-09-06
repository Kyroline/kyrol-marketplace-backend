package model

import (
	"time"

	"gorm.io/gorm"
)

type Cart struct {
	UserID    string `gorm:"type:varchar(15)"`
	ID        string `gorm:"primaryKey;type:varchar(15)"`
	ProductID string `gorm:"type:varchar(15)"`
	Qty       uint
	User      User    `gorm:"foreignKey:UserID"`
	Product   Product `gorm:"foreignKey:ProductID"`
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
