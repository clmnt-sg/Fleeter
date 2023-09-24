package config

import (
	"app/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() error {
	dsn := "host=database user=postgres password=password dbname=fleeter port=5432 sslmode=disable"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	err = database.AutoMigrate(&models.User{}, &models.Fleet{})
	if err != nil {
		panic(err)
	}
	DB = database
	return nil
}
