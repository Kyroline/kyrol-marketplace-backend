package model

import (
	"html"
	"strings"
	"time"

	"gorm.io/gorm"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        string `gorm:"primaryKey;type:varchar(15);not null"`
	FirstName string `gorm:"type:varchar(255);not null"`
	LastName  string `gorm:"type:varchar(255);not null"`
	Username  string `gorm:"type:varchar(255);unique;not null`
	Email     string `gorm:"type:varchar(255);unique;not null"`
	Password  string `gorm:"type:varchar(255);not null"`
	Roles     []Role `gorm:"many2many:user_role"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (entity *User) BeforeCreate(db *gorm.DB) error {
	entity.CreatedAt = time.Now().Local()
	return nil
}

func (entity *User) BeforeUpdate(db *gorm.DB) error {
	entity.UpdatedAt = time.Now().Local()
	return nil
}

func (entity *User) BeforeSave(db *gorm.DB) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(entity.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	entity.Password = string(hashedPassword)
	entity.Username = html.EscapeString(strings.TrimSpace(entity.Username))
	return nil
}
