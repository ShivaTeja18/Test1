package dbc

import (
	"EMP_API/Details"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func Connect() *gorm.DB {
	const DNS = `postgresql://postgres:Shiva0709@localhost:5432/postgres`
	Database, err := gorm.Open(postgres.Open(DNS), &gorm.Config{})

	if err != nil {
		log.Println(err)
	}
	if err == nil {
		log.Print("Connected")
	}
	_ = Database.AutoMigrate(&Details.EMP{})

	return Database
}
