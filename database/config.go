package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	model "kyrol-marketplace-backend/models"
	util "kyrol-marketplace-backend/utils"
)

var DB *gorm.DB

func Conn() error {
	var err error
	dsn := util.GetEnv("DB_USERNAME") + ":" + util.GetEnv("DB_PASSWORD") + "@tcp(" + util.GetEnv("DB_HOST") + ":" + util.GetEnv("DB_PORT") + ")/" + util.GetEnv("DB_DATABASE") + "?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	return nil
}

func CloseConn() {
	dbInstance, _ := DB.DB()
	_ = dbInstance.Close()
}

func Migrate() {
	DB.Debug().AutoMigrate(
		&model.Product{},
		&model.User{},
		&model.Category{},
	)
}
