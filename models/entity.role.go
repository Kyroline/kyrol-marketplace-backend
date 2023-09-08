package model

type Role struct {
	ID          uint `gorm:"primaryKey`
	Name        string
	Permissions []Permission `gorm:"many2many:role_permission"`
}
