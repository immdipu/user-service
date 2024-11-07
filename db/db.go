package db

import (
	"github.com/immdipu/user-service/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(connectionString string) error {

	var err error

	DB, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})

	if err != nil {
		return err
	}

	DB.AutoMigrate(&models.User{})

	return nil

}
