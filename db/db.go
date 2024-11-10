package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() error {

	var err error

	dns := "host=postgres_db user=postgres password=dipu dbname=user_service port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	DB, err = gorm.Open(postgres.Open(dns), &gorm.Config{})

	if err != nil {
		return err
	}

	return nil

}
