package model

import "gorm.io/gorm"

type Permission struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
