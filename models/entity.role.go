package model

import "gorm.io/gorm"

type Role struct {
	ID          uint `gorm:"primaryKey`
	Name        string
	Permissions []Permission   `gorm:"many2many:role_permission"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
