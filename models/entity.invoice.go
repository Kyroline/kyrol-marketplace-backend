package model

import (
	"time"

	"gorm.io/gorm"
)

type Invoice struct {
	ID        string        `gorm:"type:varchar(30);primaryKey;not null"`
	UserID    string        `gorm:"type:varchar(15)"`
	User      User          `gorm:"foreignKey:UserID"`
	Items     []InvoiceItem `gorm:"foreignKey:InvoiceID"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (entity *Invoice) BeforeCreate(db *gorm.DB) error {
	entity.CreatedAt = time.Now().Local()
	return nil
}

func (entity *Invoice) BeforeUpdate(db *gorm.DB) error {
	entity.UpdatedAt = time.Now().Local()
	return nil
}
