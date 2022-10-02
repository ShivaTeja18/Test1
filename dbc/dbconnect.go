package dbc

import (
	"EMP_API/Details"
	"EMP_API/Methods"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func Connect() {
	const DNS = "host = 'localhost' port = '5432' dbname = 'API' user = 'postgres' password = 'Shiva@205101' sslmode = 'prefer'"
	Database, err := gorm.Open(postgres.Open(DNS), &gorm.Config{})

	if err != nil {
		log.Println(err)
	}
	if err == nil {
		log.Print("Connected")
	}
	_ = Database.AutoMigrate(&Details.EMP{})

	Methods.DB = Database
}
